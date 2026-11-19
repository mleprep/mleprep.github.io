//go:build ignore

// 516. Longest Palindromic Subsequence (Medium)
// https://leetcode.com/problems/longest-palindromic-subsequence/
// Run: go run 0516_longest_palindromic_subsequence.go lc.go
package main

import (
	"fmt"
	"os"
)

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([]int, n) // dp[r] = answer for s[l..r] with the current l
	for l := n - 1; l >= 0; l-- {
		dp[l] = 1
		diag := 0 // old dp[l+1][r-1]
		for r := l + 1; r < n; r++ {
			tmp := dp[r]
			if s[l] == s[r] {
				dp[r] = diag + 2
			} else {
				dp[r] = max(dp[r], dp[r-1])
			}
			diag = tmp
		}
	}
	return dp[n-1]
}

func main() {
	assertEq(longestPalindromeSubseq("bbbab"), 4)
	assertEq(longestPalindromeSubseq("cbbd"), 2)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
