//go:build ignore

// 274. H-Index (Medium)
// https://leetcode.com/problems/h-index/
// Run: go run 0274_h_index.go lc.go
package main

import (
	"fmt"
	"os"
)

func hIndex(citations []int) int {
	n := len(citations)
	cnt := make([]int, n+1) // cnt[c] = papers with c citations (capped at n)
	for _, c := range citations {
		cnt[min(c, n)]++
	}
	papers := 0
	for h := n; h >= 0; h-- {
		papers += cnt[h]
		if papers >= h {
			return h
		}
	}
	return 0
}

func main() {
	assertEq(hIndex([]int{3, 0, 6, 1, 5}), 3)
	assertEq(hIndex([]int{1, 3, 1}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
