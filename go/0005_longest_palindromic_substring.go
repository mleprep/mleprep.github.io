//go:build ignore

// 5. Longest Palindromic Substring (Medium)
// https://leetcode.com/problems/longest-palindromic-substring/
// Run: go run 0005_longest_palindromic_substring.go lc.go
package main

import (
	"fmt"
	"os"
)

func longestPalindrome(s string) string {
	lo, hi := 0, 0                    // best is s[lo:hi]
	for c := 0; c < 2*len(s)-1; c++ { // 2n-1 centres: characters and gaps
		l, r := c/2, (c+1)/2
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l, r = l-1, r+1
		}
		if r-l-1 > hi-lo {
			lo, hi = l+1, r
		}
	}
	return s[lo:hi]
}

func main() {
	got := longestPalindrome("babad")
	assertEq(got == "bab" || got == "aba", true)
	assertEq(longestPalindrome("cbbd"), "bb")
	assertEq(longestPalindrome("a"), "a")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
