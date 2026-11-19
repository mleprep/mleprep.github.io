//go:build ignore

// 30. Substring with Concatenation of All Words (Hard)
// https://leetcode.com/problems/substring-with-concatenation-of-all-words/
// Run: go run 0030_substring_with_concatenation_of_all_words.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func findSubstring(s string, words []string) []int {
	L, k := len(words[0]), len(words)
	need := map[string]int{}
	for _, w := range words {
		need[w]++
	}
	res := []int{}
	for off := 0; off < L; off++ { // one sliding window per alignment
		l, used := off, 0
		cnt := map[string]int{}
		for r := off; r+L <= len(s); r += L {
			w := s[r : r+L]
			if need[w] == 0 {
				l, used, cnt = r+L, 0, map[string]int{}
				continue
			}
			cnt[w]++
			used++
			for cnt[w] > need[w] {
				cnt[s[l:l+L]]--
				used--
				l += L
			}
			if used == k {
				res = append(res, l)
			}
		}
	}
	sort.Ints(res)
	return res
}

func main() {
	assertUnordered(findSubstring("barfoothefoobarman", []string{"foo", "bar"}), []int{0, 9})
	assertEq(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}), []int{})
	assertUnordered(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}), []int{6, 9, 12})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
