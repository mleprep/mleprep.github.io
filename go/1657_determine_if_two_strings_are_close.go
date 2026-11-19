//go:build ignore

// 1657. Determine if Two Strings Are Close (Medium)
// https://leetcode.com/problems/determine-if-two-strings-are-close/
// Run: go run 1657_determine_if_two_strings_are_close.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	var c1, c2 [26]int
	for _, c := range word1 {
		c1[c-'a']++
	}
	for _, c := range word2 {
		c2[c-'a']++
	}
	for i := range c1 {
		if (c1[i] == 0) != (c2[i] == 0) { // same set of letters
			return false
		}
	}
	a, b := c1[:], c2[:]
	sort.Ints(a)
	sort.Ints(b)
	return c1 == c2 // same multiset of counts (arrays sorted in place)
}

func main() {
	assertEq(closeStrings("abc", "bca"), true)
	assertEq(closeStrings("a", "aa"), false)
	assertEq(closeStrings("cabbba", "abbccc"), true)
	assertEq(closeStrings("uau", "ssx"), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
