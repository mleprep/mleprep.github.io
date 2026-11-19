//go:build ignore

// 392. Is Subsequence (Easy)
// https://leetcode.com/problems/is-subsequence/
// Run: go run 0392_is_subsequence.go lc.go
package main

import (
	"fmt"
	"os"
)

func isSubsequence(s string, t string) bool {
	i := 0
	for j := 0; j < len(t) && i < len(s); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}

func main() {
	assertEq(isSubsequence("abc", "ahbgdc"), true)
	assertEq(isSubsequence("axc", "ahbgdc"), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
