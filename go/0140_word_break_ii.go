//go:build ignore

// 140. Word Break II (Hard)
// https://leetcode.com/problems/word-break-ii/
// Run: go run 0140_word_break_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func wordBreak(s string, wordDict []string) []string {
	words := map[string]bool{}
	for _, w := range wordDict {
		words[w] = true
	}
	memo := map[int][]string{}
	var gen func(i int) []string // all sentences for s[i:]
	gen = func(i int) []string {
		if i == len(s) {
			return []string{""}
		}
		if out, ok := memo[i]; ok {
			return out
		}
		out := []string{}
		for j := i + 1; j <= len(s); j++ {
			if w := s[i:j]; words[w] {
				for _, rest := range gen(j) {
					if rest == "" {
						out = append(out, w)
					} else {
						out = append(out, w+" "+rest)
					}
				}
			}
		}
		memo[i] = out
		return out
	}
	return gen(0)
}

func main() {
	assertUnordered(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}), []string{"cat sand dog", "cats and dog"})
	assertUnordered(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}), []string{
		"pine apple pen apple", "pine applepen apple", "pineapple pen apple"})
	assertEq(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}), []string{})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
