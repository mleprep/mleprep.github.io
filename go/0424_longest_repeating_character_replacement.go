//go:build ignore

// 424. Longest Repeating Character Replacement (Medium)
// https://leetcode.com/problems/longest-repeating-character-replacement/
// Run: go run 0424_longest_repeating_character_replacement.go lc.go
package main

import (
	"fmt"
	"os"
)

func characterReplacement(s string, k int) int {
	var cnt [26]int
	l, maxf := 0, 0
	for r := 0; r < len(s); r++ {
		cnt[s[r]-'A']++
		maxf = max(maxf, cnt[s[r]-'A'])
		if r-l+1-maxf > k { // the window never shrinks, it only slides
			cnt[s[l]-'A']--
			l++
		}
	}
	return len(s) - l
}

func main() {
	assertEq(characterReplacement("ABAB", 2), 4)
	assertEq(characterReplacement("AABABBA", 1), 4)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
