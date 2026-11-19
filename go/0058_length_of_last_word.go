//go:build ignore

// 58. Length of Last Word (Easy)
// https://leetcode.com/problems/length-of-last-word/
// Run: go run 0058_length_of_last_word.go lc.go
package main

import (
	"fmt"
	"os"
)

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	j := i
	for j >= 0 && s[j] != ' ' {
		j--
	}
	return i - j
}

func main() {
	assertEq(lengthOfLastWord("Hello World"), 5)
	assertEq(lengthOfLastWord("   fly me   to   the moon  "), 4)
	assertEq(lengthOfLastWord("luffy is still joyboy"), 6)
	assertEq(lengthOfLastWord("a"), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
