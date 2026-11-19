//go:build ignore

// 2707. Extra Characters in a String (Medium)
// https://leetcode.com/problems/extra-characters-in-a-string/
// Run: go run 2707_extra_characters_in_a_string.go lc.go
package main

import (
	"fmt"
	"os"
)

func minExtraChar(s string, dictionary []string) int {
	words, n := map[string]bool{}, len(s)
	for _, w := range dictionary {
		words[w] = true
	}
	dp := make([]int, n+1) // dp[i] = fewest extra characters in s[i:]
	for i := n - 1; i >= 0; i-- {
		dp[i] = 1 + dp[i+1] // s[i] is extra
		for j := i + 1; j <= n; j++ {
			if words[s[i:j]] {
				dp[i] = min(dp[i], dp[j])
			}
		}
	}
	return dp[0]
}

func main() {
	assertEq(minExtraChar("leetscode", []string{"leet", "code", "leetcode"}), 1)
	assertEq(minExtraChar("sayhelloworld", []string{"hello", "world"}), 3)
	assertEq(minExtraChar("abc", []string{"d"}), 3)
	assertEq(minExtraChar("abc", []string{"abc", "a"}), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
