//go:build ignore

// 336. Palindrome Pairs (Hard)
// https://leetcode.com/problems/palindrome-pairs/
// Run: go run 0336_palindrome_pairs.go lc.go
package main

import (
	"fmt"
	"os"
)

func palindromePairs(words []string) [][]int {
	isPal := func(s string) bool {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}
	reverse := func(s string) string {
		b := []byte(s)
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
		return string(b)
	}
	rev := map[string]int{}
	for i, w := range words {
		rev[reverse(w)] = i
	}
	res := [][]int{}
	for i, w := range words {
		for j := 0; j <= len(w); j++ {
			pre, suf := w[:j], w[j:]
			// w + words[k] is a palindrome if words[k] reversed == pre and suf is a palindrome
			if k, ok := rev[pre]; ok && k != i && isPal(suf) {
				res = append(res, []int{i, k})
			}
			// words[k] + w is a palindrome if words[k] reversed == suf and pre is a palindrome
			if k, ok := rev[suf]; ok && j > 0 && k != i && isPal(pre) {
				res = append(res, []int{k, i})
			}
		}
	}
	return res
}

func main() {
	assertUnordered(palindromePairs([]string{"abcd", "dcba", "lls", "s", "sssll"}), [][]int{{0, 1}, {1, 0}, {2, 4}, {3, 2}})
	assertUnordered(palindromePairs([]string{"bat", "tab", "cat"}), [][]int{{0, 1}, {1, 0}})
	assertUnordered(palindromePairs([]string{"a", ""}), [][]int{{0, 1}, {1, 0}})
	assertEq(palindromePairs([]string{"a", "b", "c"}), [][]int{})
	assertUnordered(palindromePairs([]string{"ab", "ba", "aba"}), [][]int{{0, 1}, {0, 2}, {1, 0}, {2, 1}}) // "ababa" twice
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
