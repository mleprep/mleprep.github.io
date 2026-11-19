//go:build ignore

// 583. Delete Operation for Two Strings (Medium)
// https://leetcode.com/problems/delete-operation-for-two-strings/
// Run: go run 0583_delete_operation_for_two_strings.go lc.go
package main

import (
	"fmt"
	"os"
)

func minDistance(word1 string, word2 string) int {
	n := len(word2)
	dp := make([]int, n+1) // LCS row; everything outside the LCS gets deleted
	for i := 0; i < len(word1); i++ {
		diag := 0
		for j := 1; j <= n; j++ {
			up := dp[j]
			if word1[i] == word2[j-1] {
				dp[j] = diag + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			diag = up
		}
	}
	return len(word1) + n - 2*dp[n]
}

func main() {
	assertEq(minDistance("sea", "eat"), 2)
	assertEq(minDistance("leetcode", "etco"), 4)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
