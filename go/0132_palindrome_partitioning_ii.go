//go:build ignore

// 132. Palindrome Partitioning II (Hard)
// https://leetcode.com/problems/palindrome-partitioning-ii/
// Run: go run 0132_palindrome_partitioning_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func minCut(s string) int {
	n := len(s)
	cuts := make([]int, n+1) // cuts[i] = min cuts for s[:i]; cuts[0] = -1
	for i := range cuts {
		cuts[i] = i - 1 // worst case: cut between every character
	}
	for c := 0; c < n; c++ {
		for _, lr := range [][2]int{{c, c}, {c, c + 1}} { // odd and even palindromes centred at c
			l, r := lr[0], lr[1]
			for l >= 0 && r < n && s[l] == s[r] {
				cuts[r+1] = min(cuts[r+1], cuts[l]+1)
				l, r = l-1, r+1
			}
		}
	}
	return cuts[n]
}

func main() {
	assertEq(minCut("aab"), 1)
	assertEq(minCut("a"), 0)
	assertEq(minCut("abccbc"), 2)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
