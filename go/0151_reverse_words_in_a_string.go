//go:build ignore

// 151. Reverse Words in a String (Medium)
// https://leetcode.com/problems/reverse-words-in-a-string/
// Run: go run 0151_reverse_words_in_a_string.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	assertEq(reverseWords("the sky is blue"), "blue is sky the")
	assertEq(reverseWords("  hello world  "), "world hello")
	assertEq(reverseWords("a good   example"), "example good a")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
