// Command gojudge is the site's in-browser Go judge.
//
// Built for the browser (GOOS=js GOARCH=wasm, see build.sh) it defines a global JavaScript function
//
//	goJudge(lcSource, userCode, testCode) -> JSON string
//
// that assets/go-worker.js calls. Built natively it self-checks every problem (see check.go):
//
//	go run . -data ../../data/dsa -lc ../../go/lc.go
//
// Go has no compiler that runs in a browser, so the code is interpreted by Yaegi
// (github.com/traefik/yaegi), a Go interpreter written in Go, compiled to WebAssembly.
//
// judge() mirrors assets/judge.py: the test code is the body of main(); every top-level statement that
// calls assertEq / assertUnordered / assertUnorderedDeep is one test case, run on its own so a panic
// in one case does not stop the others. Other statements (c := Constructor(2), c.Put(1, 1)) are setup:
// if one panics the remaining statements cannot run. Like LeetCode, user code needs no package clause
// or imports: packages it uses (sort, strings, heap...) are imported automatically.
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/traefik/yaegi/interp"

	"gojudge/symbols"
)

// knownPkgs: package name -> import path for packages code may use without importing them.
var knownPkgs = map[string]string{
	"bits": "math/bits", "bytes": "bytes", "cmp": "cmp", "errors": "errors", "fmt": "fmt",
	"heap": "container/heap", "list": "container/list", "maps": "maps", "math": "math",
	"rand": "math/rand", "regexp": "regexp", "slices": "slices", "sort": "sort", "strconv": "strconv",
	"strings": "strings", "unicode": "unicode", "utf8": "unicode/utf8",
}

const maxOut = 20000

// Case is one test case, in the same shape as assets/judge.py's.
type Case struct {
	Src      string  `json:"src"`
	Ok       bool    `json:"ok"`
	Got      *string `json:"got,omitempty"`
	Expected *string `json:"expected,omitempty"`
	Error    string  `json:"error,omitempty"`
	Exc      bool    `json:"exc,omitempty"`
}

// Result of judge(); error is set when the code does not compile.
type Result struct {
	Error  *string `json:"error"`
	Cases  []Case  `json:"cases"`
	Stdout string  `json:"stdout"`
	Ms     float64 `json:"ms"`
}

// ---------------------------------------------------------------- parsing

type importSpec struct{ name, path string }

// unit is a parsed source file with its package clause and imports blanked out (line numbers kept).
type unit struct {
	body    string
	imports []importSpec
	uses    map[string]bool // identifiers used as X in X.Sel that resolve to nothing in the file
	decls   map[string]bool // top-level names
}

func firstToken(src string) token.Token {
	var s scanner.Scanner
	fset := token.NewFileSet()
	s.Init(fset.AddFile("", -1, len(src)), []byte(src), nil, 0)
	for {
		_, tok, _ := s.Scan()
		if tok != token.COMMENT {
			return tok
		}
	}
}

func parseUnit(name, src string) (*unit, error) {
	prefix := ""
	if firstToken(src) != token.PACKAGE {
		prefix = "package main;" // same line, so line numbers stay those of the source
	}
	src = prefix + src
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, name, src, parser.AllErrors)
	if err != nil {
		return nil, err
	}
	b := []byte(src)
	off := func(p token.Pos) int { return fset.Position(p).Offset }
	blank := func(from, to int) {
		for i := from; i < to && i < len(b); i++ {
			if b[i] != '\n' {
				b[i] = ' '
			}
		}
	}
	end := off(f.Name.End())
	if prefix != "" {
		end++ // the ';'
	}
	blank(off(f.Package), end)
	u := &unit{uses: map[string]bool{}, decls: map[string]bool{}}
	for _, d := range f.Decls {
		switch d := d.(type) {
		case *ast.GenDecl:
			if d.Tok == token.IMPORT {
				blank(off(d.Pos()), off(d.End()))
				for _, s := range d.Specs {
					s := s.(*ast.ImportSpec)
					p, _ := strconv.Unquote(s.Path.Value)
					spec := importSpec{path: p}
					if s.Name != nil {
						spec.name = s.Name.Name
					}
					u.imports = append(u.imports, spec)
				}
				continue
			}
			for _, s := range d.Specs {
				switch s := s.(type) {
				case *ast.TypeSpec:
					u.decls[s.Name.Name] = true
				case *ast.ValueSpec:
					for _, n := range s.Names {
						u.decls[n.Name] = true
					}
				}
			}
		case *ast.FuncDecl:
			if d.Recv == nil {
				u.decls[d.Name.Name] = true
			}
		}
	}
	ast.Inspect(f, func(n ast.Node) bool {
		if sel, ok := n.(*ast.SelectorExpr); ok {
			if id, ok := sel.X.(*ast.Ident); ok && id.Obj == nil {
				u.uses[id.Name] = true
			}
		}
		return true
	})
	u.body = string(b)
	return u, nil
}

// tests is the test code split into top-level statements.
type tests struct {
	stmts  []string
	lines  []int
	isCase []bool
	uses   map[string]bool
}

var asserts = map[string]bool{"assertEq": true, "assertUnordered": true, "assertUnorderedDeep": true}

func splitTests(test string) (*tests, error) {
	const pre = "package main;func main(){\n"
	src := pre + test + "\n}"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "test.go", src, parser.AllErrors)
	if err != nil {
		return nil, err
	}
	t := &tests{uses: map[string]bool{}}
	assertish := map[string]bool{} // test-local closures that assert
	callsAssert := func(n ast.Node) bool {
		found := false
		ast.Inspect(n, func(n ast.Node) bool {
			if c, ok := n.(*ast.CallExpr); ok {
				if id, ok := c.Fun.(*ast.Ident); ok && (asserts[id.Name] || assertish[id.Name]) {
					found = true
				}
			}
			return !found
		})
		return found
	}
	body := f.Decls[len(f.Decls)-1].(*ast.FuncDecl).Body
	for _, s := range body.List {
		decl := false
		switch s := s.(type) {
		case *ast.DeclStmt:
			decl = true
		case *ast.AssignStmt:
			decl = s.Tok == token.DEFINE
		}
		has := callsAssert(s)
		if decl && has { // f := func(...) { assertEq(...) }: calls to f are test cases
			if a, ok := s.(*ast.AssignStmt); ok {
				for _, l := range a.Lhs {
					if id, ok := l.(*ast.Ident); ok {
						assertish[id.Name] = true
					}
				}
			}
		}
		p, e := fset.Position(s.Pos()), fset.Position(s.End())
		t.stmts = append(t.stmts, src[p.Offset:e.Offset])
		t.lines = append(t.lines, p.Line-1)
		t.isCase = append(t.isCase, has && !decl)
	}
	ast.Inspect(f, func(n ast.Node) bool {
		if sel, ok := n.(*ast.SelectorExpr); ok {
			if id, ok := sel.X.(*ast.Ident); ok && id.Obj == nil {
				t.uses[id.Name] = true
			}
		}
		return true
	})
	return t, nil
}

// Positions in messages: "solution.go:3:5: msg" from go/parser, "12:5: msg" (a line of the assembled
// program, see lineMap) from Yaegi.
var (
	errPos  = regexp.MustCompile(`^(?:.*?)(solution|test|lc)\.go:(\d+)(?::\d+)?:\s*`)
	progPos = regexp.MustCompile(`^(\d+):\d+:\s*`)
)

// lineMap maps lines of the assembled program back to the user's code, the test or lc.go.
// (Yaegi scopes imports by file name, so //line directives cannot be used for this.)
type lineMap []struct {
	start, end int // program lines [start, end)
	file       string
	orig       int // line of start in file
}

func (m lineMap) locate(line int) (string, int) {
	for _, r := range m {
		if line >= r.start && line < r.end {
			return r.file, r.orig + line - r.start
		}
	}
	return "", 0
}

func where(file, line string) string {
	switch file {
	case "solution":
		return " (line " + line + " of your code)"
	case "test":
		return " (line " + line + " of test)"
	}
	return " (line " + line + " of lc.go)"
}

// errLine rewrites a positioned error as "msg (line 3 of your code)".
func errLine(err error, lm lineMap) string {
	var list scanner.ErrorList
	if errors.As(err, &list) && len(list) > 0 {
		err = list[0]
	}
	msg := err.Error()
	if m := errPos.FindStringSubmatch(msg); m != nil {
		return strings.TrimSpace(msg[len(m[0]):]) + where(m[1], m[2])
	}
	if m := progPos.FindStringSubmatch(msg); m != nil {
		n, _ := strconv.Atoi(m[1])
		if file, line := lm.locate(n); file != "" {
			return strings.TrimSpace(msg[len(m[0]):]) + where(file, strconv.Itoa(line))
		}
		return strings.TrimSpace(msg[len(m[0]):])
	}
	return msg
}

// ---------------------------------------------------------------- min / max

// Yaegi predates Go 1.21's min and max builtins; unless the code defines its own, these stand in.
const minMax = `
func min[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string](x T, y ...T) T {
	for _, v := range y {
		if v < x {
			x = v
		}
	}
	return x
}

func max[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string](x T, y ...T) T {
	for _, v := range y {
		if v > x {
			x = v
		}
	}
	return x
}
`

// ---------------------------------------------------------------- running

type harness struct {
	t      *tests
	lm     lineMap
	cases  []Case
	cur    *Case
	stmt   int
	stderr *bytes.Buffer
}

// Yaegi's panic trace has one "line:col: panic: main.f(...)" line per interpreted frame, innermost
// first; the position is that of the function's first statement, so name the function instead.
var panicPos = regexp.MustCompile(`(?m)^(\d+):\d+: panic: (.*)\(\.\.\.\)\s*$`)

// panicWhere names the innermost function of the user's code (else the test line) that was running.
func (h *harness) panicWhere() string {
	ms := panicPos.FindAllStringSubmatch(h.stderr.String(), -1)
	for _, m := range ms {
		n, _ := strconv.Atoi(m[1])
		if file, _ := h.lm.locate(n); file == "solution" {
			name := strings.TrimPrefix(m[2], "main.")
			if i := strings.Index(name, ".func"); i >= 0 {
				return " (in a function literal inside " + name[:i] + "(), your code)"
			}
			return " (in " + name + "(), your code)"
		}
	}
	for _, m := range ms {
		n, _ := strconv.Atoi(m[1])
		if file, line := h.lm.locate(n); file == "test" {
			return where(file, strconv.Itoa(line))
		}
	}
	return ""
}

// panicText words Yaegi's (reflect-based) runtime panics like the Go runtime's.
func panicText(v any) string {
	if p, ok := v.(interp.Panic); ok {
		v = p.Value
	}
	s := fmt.Sprint(v)
	if _, ok := v.(runtime.Error); ok && !strings.HasPrefix(s, "reflect") {
		return s
	}
	switch {
	case strings.Contains(s, "index out of range"):
		return "runtime error: index out of range"
	case strings.Contains(s, "slice bounds out of range"):
		return "runtime error: slice bounds out of range"
	case strings.Contains(s, "zero Value") || strings.Contains(s, "nil pointer"):
		return "runtime error: invalid memory address or nil pointer dereference"
	case strings.Contains(s, "nil map"):
		return "assignment to entry in nil map"
	case strings.Contains(s, "integer divide by zero"):
		return "runtime error: integer divide by zero"
	}
	return "panic: " + s
}

func (h *harness) runCase(i int, f func()) {
	h.stmt = i
	h.stderr.Reset()
	c := Case{Src: h.t.stmts[i], Ok: true}
	h.cur = &c
	defer func() {
		if v := recover(); v != nil {
			c.Ok, c.Exc, c.Got, c.Expected = false, true, nil, nil
			c.Error = panicText(v) + h.panicWhere()
		}
		h.cur = nil
		h.cases = append(h.cases, c)
	}()
	f()
}

func (h *harness) report(ok bool, got, want string) {
	c := h.cur
	if c == nil { // an assertion inside a setup statement: count it as a case of its own
		h.cases = append(h.cases, Case{Src: h.t.stmts[h.stmt], Ok: true})
		c = &h.cases[len(h.cases)-1]
	}
	if !ok && c.Ok {
		c.Ok, c.Got, c.Expected = false, &got, &want
	}
}

func (h *harness) exports() interp.Exports {
	return interp.Exports{"lcjudge/lcjudge": {
		"Case":   reflect.ValueOf(h.runCase),
		"Stmt":   reflect.ValueOf(func(i int) { h.stmt = i; h.stderr.Reset() }),
		"Report": reflect.ValueOf(h.report),
	}}
}

var judgeMu sync.Mutex

var lcCache struct {
	sync.Mutex
	src string
	u   *unit
}

func parseLC(src string) (*unit, error) {
	lcCache.Lock()
	defer lcCache.Unlock()
	if lcCache.u == nil || lcCache.src != src {
		u, err := parseUnit("lc.go", src)
		if err != nil {
			return nil, err
		}
		lcCache.src, lcCache.u = src, u
	}
	return lcCache.u, nil
}

// program assembles the file Yaegi runs: imports, the user's code, lc.go, and main() with the tests.
func program(lc, user *unit, t *tests) (string, lineMap) {
	paths := map[string]map[string]bool{} // path -> names
	add := func(name, path string) {
		if paths[path] == nil {
			paths[path] = map[string]bool{}
		}
		paths[path][name] = true
	}
	imported := map[string]bool{}
	for _, u := range []*unit{lc, user} {
		for _, s := range u.imports {
			add(s.name, s.path)
			n := s.name
			if n == "" {
				n = s.path[strings.LastIndex(s.path, "/")+1:]
			}
			imported[n] = true
		}
	}
	for _, uses := range []map[string]bool{user.uses, t.uses} {
		for n := range uses {
			if p, ok := knownPkgs[n]; ok && !imported[n] && !user.decls[n] && !lc.decls[n] {
				add("", p)
				imported[n] = true
			}
		}
	}
	var b strings.Builder
	var lm lineMap
	line := 1
	emit := func(text, file string, orig int) {
		n := strings.Count(text, "\n")
		if file != "" {
			lm = append(lm, struct {
				start, end int
				file       string
				orig       int
			}{line, line + n + 1, file, orig})
		}
		b.WriteString(text)
		line += n
	}
	emit("package main\n\nimport (\n\tlcjudge \"lcjudge\"\n", "", 0)
	keys := make([]string, 0, len(paths))
	for p := range paths {
		keys = append(keys, p)
	}
	sort.Strings(keys)
	for _, p := range keys {
		names := make([]string, 0, len(paths[p]))
		for n := range paths[p] {
			names = append(names, n)
		}
		sort.Strings(names)
		for _, n := range names {
			emit(fmt.Sprintf("\t%s %q\n", n, p), "", 0)
		}
	}
	emit(")\n", "", 0)
	if !user.decls["min"] && !user.decls["max"] {
		emit(minMax, "", 0)
	}
	emit("\n", "", 0)
	emit(user.body, "solution", 1)
	emit("\n", "", 0)
	emit(lc.body, "lc", 1)
	emit("\nfunc main() {\n\tlcReport = lcjudge.Report\n", "", 0)
	for i, s := range t.stmts {
		if t.isCase[i] {
			emit(fmt.Sprintf("\tlcjudge.Case(%d, func() {\n", i), "", 0)
			emit(s, "test", t.lines[i])
			emit("\n\t})\n", "", 0)
		} else {
			emit(fmt.Sprintf("\tlcjudge.Stmt(%d)\n", i), "", 0)
			emit(s, "test", t.lines[i])
			emit("\n", "", 0)
		}
	}
	emit("}\n", "", 0)
	return b.String(), lm
}

func toJSON(r Result) string {
	b, _ := json.Marshal(r)
	return string(b)
}

// judge runs code against test and returns the Result as JSON.
func judge(lcSrc, code, test string) (out string) {
	// One run at a time: Yaegi builds recursive types (ListNode...) by patching reflect's globally cached
	// struct types, so two interpreters building the same type at once corrupt each other.
	judgeMu.Lock()
	defer judgeMu.Unlock()
	t0 := time.Now()
	res := Result{Cases: []Case{}}
	defer func() { // the interpreter itself can panic on code it does not support
		if v := recover(); v != nil {
			if os.Getenv("GOJUDGE_DEBUG") != "" {
				debug.PrintStack()
			}
			msg := fmt.Sprintf("the Go interpreter (Yaegi) failed on this code: %v", v)
			res.Error, res.Cases = &msg, []Case{}
			out = toJSON(res)
		}
	}()
	done := func() string {
		res.Ms = float64(time.Since(t0).Microseconds()/100) / 10
		if len(res.Stdout) > maxOut {
			res.Stdout = res.Stdout[:maxOut] + fmt.Sprintf("\n… (%d more characters)", len(res.Stdout)-maxOut)
		}
		return toJSON(res)
	}
	fail := func(msg string) string { res.Error = &msg; return done() }

	lc, err := parseLC(lcSrc)
	if err != nil {
		return fail("lc.go: " + err.Error())
	}
	user, err := parseUnit("solution.go", code)
	if err != nil {
		return fail("syntax error: " + errLine(err, nil))
	}
	t, err := splitTests(test)
	if err != nil {
		return fail("test parse error: " + err.Error())
	}
	src, lm := program(lc, user, t)

	var stdout, stderr bytes.Buffer
	i := interp.New(interp.Options{Stdout: &stdout, Stderr: &stderr})
	if err := i.Use(symbols.Symbols); err != nil {
		return fail("judge: " + err.Error())
	}
	h := &harness{t: t, lm: lm, stderr: &stderr}
	if err := i.Use(h.exports()); err != nil {
		return fail("judge: " + err.Error())
	}
	prog, err := i.Compile(src)
	if err != nil {
		return fail("compile error: " + errLine(err, lm))
	}
	t0 = time.Now() // time the run, not the compilation
	_, err = i.Execute(prog)
	res.Stdout = stdout.String()
	if err != nil {
		var p interp.Panic
		if errors.As(err, &p) {
			// a setup statement panicked: report it, the statements after it could not run
			res.Cases = append(h.cases, Case{Src: t.stmts[h.stmt], Ok: false, Exc: true, Error: panicText(p.Value) + h.panicWhere()})
			return done()
		}
		return fail(errLine(err, lm))
	}
	res.Cases = h.cases
	return done()
}
