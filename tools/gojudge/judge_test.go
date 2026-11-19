package main

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

// Regression tests for the judge itself (go test in tools/gojudge). Problem solutions are checked by
// `go run .` (the self-check) instead.

func runJudge(t *testing.T, code, test string) Result {
	t.Helper()
	lc, err := os.ReadFile("../../go/lc.go")
	if err != nil {
		t.Fatal(err)
	}
	var r Result
	if err := json.Unmarshal([]byte(judge(string(lc), code, test)), &r); err != nil {
		t.Fatal(err)
	}
	return r
}

func TestJudge(t *testing.T) {
	cases := []struct {
		name, code, test string
		want             []bool // per test case: passed?
		errHas           string // substring of the compile error, or of the first failing case's error
	}{
		{name: "pass and wrong answer", code: "func f(x int) int { return x * 2 }",
			test: "assertEq(f(2), 4)\nassertEq(f(3), 7)", want: []bool{true, false}},
		{name: "auto imports, min/max, slices", code: "func f(a []int) int {\n\tsort.Ints(a)\n\treturn max(a[0], 1, 2) + slices.Max(a)\n}",
			test: "assertEq(f([]int{3, 1}), 5)", want: []bool{true}},
		{name: "multiple assignment with composite literal (patched yaegi)", code: "func f(p string) []int {\n\tk, res := len(p), []int{}\n\treturn append(res, k)\n}",
			test: "assertEq(f(\"abc\"), []int{3})", want: []bool{true}},
		{name: "panic in one case does not stop the others", code: "func f(a []int) int { return a[5] }",
			test: "assertEq(f(nil), 0)\nassertEq(1, 1)", want: []bool{false, true}, errHas: "index out of range (in f(), your code)"},
		{name: "setup statement panics", code: "type T struct{ m map[int]int }\n\nfunc (t *T) Put(k int) { t.m[k] = 1 }",
			test: "x := T{}\nx.Put(1)\nassertEq(1, 1)", want: []bool{false}, errHas: "nil map"},
		{name: "compile error line", code: "func f() int {\n\treturn y\n}",
			test: "assertEq(f(), 1)", errHas: "undefined: y (line 2 of your code)"},
		{name: "syntax error line", code: "func f() int {\n\treturn (\n}",
			test: "assertEq(f(), 1)", errHas: "line 3 of your code"},
		{name: "lists, trees and unordered", code: "func f() *TreeNode { return makeTree(1, nil, 2) }",
			test: "assertEq(f(), []any{1, nil, 2})\nassertEq(makeList(1, 2), []int{1, 2})\nassertUnorderedDeep([][]int{{2, 1}, {3}}, [][]int{{3}, {1, 2}})",
			want: []bool{true, true, true}},
		{name: "test-local closure with asserts", code: "func f(x int) int { return x }",
			test: "chk := func(x int) { assertEq(f(x), x) }\nchk(1)\nchk(2)", want: []bool{true, true}},
		{name: "stdout", code: "func f() int { fmt.Println(\"hi\"); return 1 }", test: "assertEq(f(), 1)", want: []bool{true}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := runJudge(t, c.code, c.test)
			if c.want == nil {
				if r.Error == nil || !strings.Contains(*r.Error, c.errHas) {
					t.Fatalf("want error containing %q, got %+v", c.errHas, r)
				}
				return
			}
			if r.Error != nil {
				t.Fatalf("unexpected error %s", *r.Error)
			}
			if len(r.Cases) != len(c.want) {
				t.Fatalf("want %d cases, got %+v", len(c.want), r.Cases)
			}
			for i, ok := range c.want {
				if r.Cases[i].Ok != ok {
					t.Fatalf("case %d: want ok=%v, got %+v", i, ok, r.Cases[i])
				}
			}
			if c.errHas != "" {
				for _, cs := range r.Cases {
					if !cs.Ok {
						if !strings.Contains(cs.Error, c.errHas) {
							t.Fatalf("want error containing %q, got %q", c.errHas, cs.Error)
						}
						break
					}
				}
			}
			if c.name == "stdout" && r.Stdout != "hi\n" {
				t.Fatalf("stdout %q", r.Stdout)
			}
		})
	}
}
