//go:build ignore

// 97. Interleaving String (Medium)
// https://leetcode.com/problems/interleaving-string/
// Run: go run 0097_interleaving_string.go lc.go
package main

import (
	"fmt"
	"os"
)

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	dp := make([]bool, n+1) // dp[j]: s1[:i] and s2[:j] interleave to s3[:i+j]
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				dp[j] = true
				continue
			}
			fromS1 := i > 0 && dp[j] && s1[i-1] == s3[i+j-1]
			fromS2 := j > 0 && dp[j-1] && s2[j-1] == s3[i+j-1]
			dp[j] = fromS1 || fromS2
		}
	}
	return dp[n]
}

func main() {
	assertEq(isInterleave("aabcc", "dbbca", "aadbbcbcac"), true)
	assertEq(isInterleave("aabcc", "dbbca", "aadbbbaccc"), false)
	assertEq(isInterleave("", "", ""), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
