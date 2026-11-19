//go:build ignore

// 1456. Maximum Number of Vowels in a Substring of Given Length (Medium)
// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
// Run: go run 1456_maximum_number_of_vowels_in_a_substring_of_given_length.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func maxVowels(s string, k int) int {
	vowel := func(c byte) int {
		if strings.IndexByte("aeiou", c) >= 0 {
			return 1
		}
		return 0
	}
	cur := 0
	for i := 0; i < k; i++ {
		cur += vowel(s[i])
	}
	best := cur
	for i := k; i < len(s); i++ {
		cur += vowel(s[i]) - vowel(s[i-k])
		best = max(best, cur)
	}
	return best
}

func main() {
	assertEq(maxVowels("abciiidef", 3), 3)
	assertEq(maxVowels("aeiou", 2), 2)
	assertEq(maxVowels("leetcode", 3), 2)
	assertEq(maxVowels("tryhard", 4), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
