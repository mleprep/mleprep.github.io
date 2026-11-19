//go:build ignore

// 3. Longest Substring Without Repeating Characters (Medium)
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
// Run: go run 0003_longest_substring_without_repeating_characters.go lc.go
package main

import (
	"fmt"
	"os"
)

func lengthOfLongestSubstring(s string) int {
	last := map[byte]int{} // char -> last index
	l, best := 0, 0
	for r := 0; r < len(s); r++ {
		if i, ok := last[s[r]]; ok && i >= l {
			l = i + 1 // jump past the previous copy
		}
		last[s[r]] = r
		best = max(best, r-l+1)
	}
	return best
}

func main() {
	assertEq(lengthOfLongestSubstring("abcabcbb"), 3)
	assertEq(lengthOfLongestSubstring("bbbbb"), 1)
	assertEq(lengthOfLongestSubstring("pwwkew"), 3)
	assertEq(lengthOfLongestSubstring("abba"), 2)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
