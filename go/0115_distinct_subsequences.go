//go:build ignore

// 115. Distinct Subsequences (Hard)
// https://leetcode.com/problems/distinct-subsequences/
// Run: go run 0115_distinct_subsequences.go lc.go
package main

import (
	"fmt"
	"os"
)

func numDistinct(s string, t string) int {
	dp := make([]int, len(t)+1) // dp[j] = ways to form t[:j] from the prefix of s seen so far
	dp[0] = 1
	for i := 0; i < len(s); i++ {
		for j := len(t); j > 0; j-- { // backwards so s[i] is used once per subsequence
			if t[j-1] == s[i] {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[len(t)]
}

func main() {
	assertEq(numDistinct("rabbbit", "rabbit"), 3)
	assertEq(numDistinct("babgbag", "bag"), 5)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
