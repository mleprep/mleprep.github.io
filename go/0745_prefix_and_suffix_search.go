//go:build ignore

// 745. Prefix and Suffix Search (Hard)
// https://leetcode.com/problems/prefix-and-suffix-search/
// Run: go run 0745_prefix_and_suffix_search.go lc.go
package main

import (
	"fmt"
	"os"
)

type WordFilter struct {
	idx map[string]int // "prefix#suffix" -> largest index
}

func Constructor(words []string) WordFilter {
	idx := map[string]int{}
	for i, w := range words {
		for p := 0; p <= len(w); p++ {
			for s := 0; s <= len(w); s++ {
				idx[w[:p]+"#"+w[s:]] = i // later words overwrite earlier ones
			}
		}
	}
	return WordFilter{idx}
}

func (this *WordFilter) F(pref string, suff string) int {
	if i, ok := this.idx[pref+"#"+suff]; ok {
		return i
	}
	return -1
}

func main() {
	wf := Constructor([]string{"apple", "ape", "maple"})
	assertEq([]int{wf.F("a", "e"), wf.F("ap", "le"), wf.F("", "le"), wf.F("b", "")}, []int{1, 0, 2, -1})
	wf2 := Constructor([]string{"apple"})
	assertEq(wf2.F("a", "e"), 0)
	wf3 := Constructor([]string{"abc", "abc"})
	assertEq(wf3.F("a", "c"), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
