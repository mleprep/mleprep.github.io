//go:build ignore

// 409. Longest Palindrome (Easy)
// https://leetcode.com/problems/longest-palindrome/
// Run: go run 0409_longest_palindrome.go lc.go
package main

import (
	"fmt"
	"os"
)

func longestPalindrome(s string) int {
	var cnt [128]int
	for i := 0; i < len(s); i++ {
		cnt[s[i]]++
	}
	pairs := 0
	for _, v := range cnt {
		pairs += v / 2 * 2
	}
	if pairs < len(s) {
		return pairs + 1 // one odd letter can sit in the middle
	}
	return pairs
}

func main() {
	assertEq(longestPalindrome("abccccdd"), 7)
	assertEq(longestPalindrome("a"), 1)
	assertEq(longestPalindrome("Aa"), 1)
	assertEq(longestPalindrome("bb"), 2)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
