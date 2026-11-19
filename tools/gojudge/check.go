//go:build !js

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type item struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Go     string `json:"go"`
	GoTest string `json:"goTest"`
	GoStub string `json:"goStub"`
}

type outcome struct {
	it       item
	ref, stb Result
	refOK    bool
	stubOK   bool
	ms       float64
}

func run(lc, code, test string, limit time.Duration) Result {
	ch := make(chan string, 1)
	go func() { ch <- judge(lc, code, test) }()
	select {
	case s := <-ch:
		var r Result
		_ = json.Unmarshal([]byte(s), &r)
		return r
	case <-time.After(limit):
		msg := "TIMEOUT"
		return Result{Error: &msg}
	}
}

func passed(r Result) bool {
	if r.Error != nil || len(r.Cases) == 0 {
		return false
	}
	for _, c := range r.Cases {
		if !c.Ok {
			return false
		}
	}
	return true
}

// Self-check: every reference solution passes its tests and every generated starter stub fails them.
func main() {
	data := flag.String("data", "../../data/dsa", "directory of topic JSON files")
	lcPath := flag.String("lc", "../../go/lc.go", "path of go/lc.go")
	only := flag.String("only", "", "comma-separated problem numbers")
	code := flag.String("code", "", "judge this file against -test and print the JSON result")
	test := flag.String("test", "", "test file for -code")
	dump := flag.Bool("dump", false, "with -code: print the program the judge would run")
	flag.Parse()
	lcb, err := os.ReadFile(*lcPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	lc := string(lcb)
	if *code != "" {
		c, _ := os.ReadFile(*code)
		t, _ := os.ReadFile(*test)
		if *dump {
			l, _ := parseUnit("lc.go", lc)
			u, err := parseUnit("solution.go", string(c))
			ts, err2 := splitTests(string(t))
			if err != nil || err2 != nil {
				fmt.Println(err, err2)
				os.Exit(2)
			}
			src, _ := program(l, u, ts)
			fmt.Print(src)
			return
		}
		fmt.Println(judge(lc, string(c), string(t)))
		return
	}
	want := map[int]bool{}
	for _, s := range strings.Split(*only, ",") {
		if n, err := strconv.Atoi(strings.TrimSpace(s)); err == nil {
			want[n] = true
		}
	}
	files, _ := filepath.Glob(filepath.Join(*data, "*.json"))
	var items []item
	total := 0
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		var topic struct{ Items []item }
		if err := json.Unmarshal(b, &topic); err != nil {
			fmt.Println(f, err)
			os.Exit(2)
		}
		for _, it := range topic.Items {
			total++
			if it.Go != "" && (len(want) == 0 || want[it.ID]) {
				items = append(items, it)
			}
		}
	}
	out := make([]outcome, len(items)) // judge() runs one program at a time, so check them in order
	for i, it := range items {
		o := outcome{it: it}
		o.ref = run(lc, it.Go, it.GoTest, 30*time.Second)
		o.stb = run(lc, it.GoStub, it.GoTest, 30*time.Second)
		o.refOK = passed(o.ref)
		o.stubOK = o.stb.Error == nil && len(o.stb.Cases) > 0 && !passed(o.stb)
		o.ms = o.ref.Ms
		out[i] = o
		if (o.ref.Error != nil && *o.ref.Error == "TIMEOUT") || (o.stb.Error != nil && *o.stb.Error == "TIMEOUT") {
			fmt.Printf("%d %s: TIMEOUT (stopping: the stuck run holds the judge)\n", it.ID, it.Title)
			os.Exit(1)
		}
	}
	bad := 0
	for _, o := range out {
		if !o.refOK || !o.stubOK {
			bad++
			fmt.Printf("%d %s:\n", o.it.ID, o.it.Title)
			if !o.refOK {
				b, _ := json.MarshalIndent(o.ref, "    ", "  ")
				fmt.Printf("    reference: %s\n", b)
			}
			if !o.stubOK {
				b, _ := json.Marshal(o.stb)
				fmt.Printf("    stub: %.600s\n", b)
			}
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ms > out[j].ms })
	fmt.Print("slowest:")
	for i := 0; i < len(out) && i < 5; i++ {
		fmt.Printf(" %d (%.0f ms)", out[i].it.ID, out[i].ms)
	}
	fmt.Printf("\ngo judge self-check: %d/%d problems have Go; ", len(items), total)
	if bad == 0 {
		fmt.Println("all good")
		return
	}
	fmt.Printf("%d problem(s) need attention\n", bad)
	os.Exit(1)
}
