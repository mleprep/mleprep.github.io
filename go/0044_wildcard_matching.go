//go:build ignore

// 44. Wildcard Matching (Hard)
// https://leetcode.com/problems/wildcard-matching/
// Run: go run 0044_wildcard_matching.go lc.go
package main

import (
	"fmt"
	"os"
)

func isMatch(s string, p string) bool {
	n := len(p)
	dp := make([]bool, n+1) // dp[j]: the current prefix of s matches p[:j]
	dp[0] = true
	for j := 1; j <= n; j++ {
		dp[j] = dp[j-1] && p[j-1] == '*'
	}
	for i := 0; i < len(s); i++ {
		diag := dp[0]
		dp[0] = false
		for j := 1; j <= n; j++ {
			up := dp[j]
			if p[j-1] == '*' {
				dp[j] = dp[j-1] || up // '*' matches nothing, or one more character
			} else {
				dp[j] = diag && (p[j-1] == s[i] || p[j-1] == '?')
			}
			diag = up
		}
	}
	return dp[n]
}

func main() {
	assertEq(isMatch("aa", "a"), false)
	assertEq(isMatch("aa", "*"), true)
	assertEq(isMatch("cb", "?a"), false)
	assertEq(isMatch("adceb", "*a*b"), true)
	assertEq(isMatch("acdcb", "a*c?b"), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
