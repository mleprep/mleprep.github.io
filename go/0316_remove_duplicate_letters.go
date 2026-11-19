//go:build ignore

// 316. Remove Duplicate Letters (Medium)
// https://leetcode.com/problems/remove-duplicate-letters/
// Run: go run 0316_remove_duplicate_letters.go lc.go
package main

import (
	"fmt"
	"os"
)

func removeDuplicateLetters(s string) string {
	var last [26]int
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}
	var used [26]bool
	st := []byte{}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if used[ch-'a'] {
			continue
		}
		// pop bigger letters that appear again later
		for len(st) > 0 && st[len(st)-1] > ch && last[st[len(st)-1]-'a'] > i {
			used[st[len(st)-1]-'a'] = false
			st = st[:len(st)-1]
		}
		st = append(st, ch)
		used[ch-'a'] = true
	}
	return string(st)
}

func main() {
	assertEq(removeDuplicateLetters("bcabc"), "abc")
	assertEq(removeDuplicateLetters("cbacdcbc"), "acdb")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
