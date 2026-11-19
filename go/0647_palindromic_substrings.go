//go:build ignore

// 647. Palindromic Substrings (Medium)
// https://leetcode.com/problems/palindromic-substrings/
// Run: go run 0647_palindromic_substrings.go lc.go
package main

import (
	"fmt"
	"os"
)

func countSubstrings(s string) int {
	count := 0
	for c := 0; c < 2*len(s)-1; c++ { // expand around each of the 2n-1 centres
		l, r := c/2, (c+1)/2
		for l >= 0 && r < len(s) && s[l] == s[r] {
			count++
			l, r = l-1, r+1
		}
	}
	return count
}

func main() {
	assertEq(countSubstrings("abc"), 3)
	assertEq(countSubstrings("aaa"), 6)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
