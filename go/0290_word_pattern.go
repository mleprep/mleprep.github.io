//go:build ignore

// 290. Word Pattern (Easy)
// https://leetcode.com/problems/word-pattern/
// Run: go run 0290_word_pattern.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)
	if len(words) != len(pattern) {
		return false
	}
	p2w, w2p := map[byte]string{}, map[string]byte{}
	for i, w := range words {
		p := pattern[i]
		if x, ok := p2w[p]; ok && x != w {
			return false
		}
		if x, ok := w2p[w]; ok && x != p {
			return false
		}
		p2w[p], w2p[w] = w, p
	}
	return true
}

func main() {
	assertEq(wordPattern("abba", "dog cat cat dog"), true)
	assertEq(wordPattern("abba", "dog cat cat fish"), false)
	assertEq(wordPattern("aaaa", "dog cat cat dog"), false)
	assertEq(wordPattern("abba", "dog dog dog dog"), false)
	assertEq(wordPattern("aaa", "aa aa aa aa"), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
