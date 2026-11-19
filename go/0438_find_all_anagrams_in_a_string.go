//go:build ignore

// 438. Find All Anagrams in a String (Medium)
// https://leetcode.com/problems/find-all-anagrams-in-a-string/
// Run: go run 0438_find_all_anagrams_in_a_string.go lc.go
package main

import (
	"fmt"
	"os"
)

func findAnagrams(s string, p string) []int {
	k, res := len(p), []int{}
	var need, win [26]int
	for i := 0; i < k; i++ {
		need[p[i]-'a']++
	}
	for r := 0; r < len(s); r++ {
		win[s[r]-'a']++
		if r >= k {
			win[s[r-k]-'a']--
		}
		if win == need {
			res = append(res, r-k+1)
		}
	}
	return res
}

func main() {
	assertUnordered(findAnagrams("cbaebabacd", "abc"), []int{0, 6})
	assertUnordered(findAnagrams("abab", "ab"), []int{0, 1, 2})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
