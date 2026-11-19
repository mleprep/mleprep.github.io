//go:build ignore

// 28. Find the Index of the First Occurrence in a String (Easy)
// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
// Run: go run 0028_find_the_index_of_the_first_occurrence_in_a_string.go lc.go
package main

import (
	"fmt"
	"os"
)

func strStr(haystack string, needle string) int {
	lps, k := make([]int, len(needle)), 0 // KMP failure function
	for i := 1; i < len(needle); i++ {
		for k > 0 && needle[i] != needle[k] {
			k = lps[k-1]
		}
		if needle[i] == needle[k] {
			k++
		}
		lps[i] = k
	}
	k = 0
	for i := 0; i < len(haystack); i++ {
		for k > 0 && haystack[i] != needle[k] {
			k = lps[k-1]
		}
		if haystack[i] == needle[k] {
			k++
			if k == len(needle) {
				return i - k + 1
			}
		}
	}
	return -1
}

func main() {
	assertEq(strStr("sadbutsad", "sad"), 0)
	assertEq(strStr("leetcode", "leeto"), -1)
	assertEq(strStr("aaaaab", "aab"), 3)
	assertEq(strStr("mississippi", "issip"), 4)
	assertEq(strStr("a", "a"), 0)
	assertEq(strStr("abc", "c"), 2)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
