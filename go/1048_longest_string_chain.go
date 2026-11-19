//go:build ignore

// 1048. Longest String Chain (Medium)
// https://leetcode.com/problems/longest-string-chain/
// Run: go run 1048_longest_string_chain.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	dp := map[string]int{} // longest chain ending at word
	best := 0
	for _, w := range words {
		dp[w] = 1
		for i := 0; i < len(w); i++ { // predecessor: w with one letter removed
			if p, ok := dp[w[:i]+w[i+1:]]; ok {
				dp[w] = max(dp[w], p+1)
			}
		}
		best = max(best, dp[w])
	}
	return best
}

func main() {
	assertEq(longestStrChain([]string{"a", "b", "ba", "bca", "bda", "bdca"}), 4)
	assertEq(longestStrChain([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}), 5)
	assertEq(longestStrChain([]string{"abcd", "dbqca"}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
