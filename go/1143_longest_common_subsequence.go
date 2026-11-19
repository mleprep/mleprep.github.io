//go:build ignore

// 1143. Longest Common Subsequence (Medium)
// https://leetcode.com/problems/longest-common-subsequence/
// Run: go run 1143_longest_common_subsequence.go lc.go
package main

import (
	"fmt"
	"os"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	n := len(text2)
	dp := make([]int, n+1) // one row of the LCS table
	for i := 0; i < len(text1); i++ {
		diag := 0
		for j := 1; j <= n; j++ {
			up := dp[j]
			if text1[i] == text2[j-1] {
				dp[j] = diag + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			diag = up
		}
	}
	return dp[n]
}

func main() {
	assertEq(longestCommonSubsequence("abcde", "ace"), 3)
	assertEq(longestCommonSubsequence("abc", "def"), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
