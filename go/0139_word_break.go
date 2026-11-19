//go:build ignore

// 139. Word Break (Medium)
// https://leetcode.com/problems/word-break/
// Run: go run 0139_word_break.go lc.go
package main

import (
	"fmt"
	"os"
)

func wordBreak(s string, wordDict []string) bool {
	words, L := map[string]bool{}, 0
	for _, w := range wordDict {
		words[w] = true
		L = max(L, len(w))
	}
	dp := make([]bool, len(s)+1) // dp[i]: s[:i] can be segmented
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := max(0, i-L); j < i; j++ { // the last word is at most L long
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func main() {
	assertEq(wordBreak("leetcode", []string{"leet", "code"}), true)
	assertEq(wordBreak("applepenapple", []string{"apple", "pen"}), true)
	assertEq(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
