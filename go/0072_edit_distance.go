//go:build ignore

// 72. Edit Distance (Medium)
// https://leetcode.com/problems/edit-distance/
// Run: go run 0072_edit_distance.go lc.go
package main

import (
	"fmt"
	"os"
)

func minDistance(word1 string, word2 string) int {
	n := len(word2)
	dp := make([]int, n+1) // dp[j] = edits turning word1[:i] into word2[:j]
	for j := range dp {
		dp[j] = j
	}
	for i := 1; i <= len(word1); i++ {
		diag := dp[0]
		dp[0] = i
		for j := 1; j <= n; j++ {
			up := dp[j]
			if word1[i-1] == word2[j-1] {
				dp[j] = diag
			} else {
				dp[j] = 1 + min(diag, up, dp[j-1]) // replace, delete, insert
			}
			diag = up
		}
	}
	return dp[n]
}

func main() {
	assertEq(minDistance("horse", "ros"), 3)
	assertEq(minDistance("intention", "execution"), 5)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
