//go:build ignore

// 1768. Merge Strings Alternately (Easy)
// https://leetcode.com/problems/merge-strings-alternately/
// Run: go run 1768_merge_strings_alternately.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	var sb strings.Builder
	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			sb.WriteByte(word1[i])
		}
		if i < len(word2) {
			sb.WriteByte(word2[i])
		}
	}
	return sb.String()
}

func main() {
	assertEq(mergeAlternately("abc", "pqr"), "apbqcr")
	assertEq(mergeAlternately("ab", "pqrs"), "apbqrs")
	assertEq(mergeAlternately("abcd", "pq"), "apbqcd")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
