//go:build ignore

// 1312. Minimum Insertion Steps to Make a String Palindrome (Hard)
// https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/
// Run: go run 1312_minimum_insertion_steps_to_make_a_string_palindrome.go lc.go
package main

import (
	"fmt"
	"os"
)

func minInsertions(s string) int {
	n := len(s)
	dp := make([]int, n) // longest palindromic subsequence, one row at a time
	for l := n - 1; l >= 0; l-- {
		dp[l] = 1
		diag := 0
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
	return n - dp[n-1] // every character outside it needs a mirror
}

func main() {
	assertEq(minInsertions("zzazz"), 0)
	assertEq(minInsertions("mbadm"), 2)
	assertEq(minInsertions("leetcode"), 5)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
