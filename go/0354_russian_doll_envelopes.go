//go:build ignore

// 354. Russian Doll Envelopes (Hard)
// https://leetcode.com/problems/russian-doll-envelopes/
// Run: go run 0354_russian_doll_envelopes.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool { // width up, and height down within a width
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	tails := []int{} // then LIS on heights
	for _, e := range envelopes {
		i := sort.SearchInts(tails, e[1])
		if i == len(tails) {
			tails = append(tails, e[1])
		} else {
			tails[i] = e[1]
		}
	}
	return len(tails)
}

func main() {
	assertEq(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}), 3)
	assertEq(maxEnvelopes([][]int{{1, 1}, {1, 1}, {1, 1}}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
