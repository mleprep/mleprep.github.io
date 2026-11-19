//go:build ignore

// 383. Ransom Note (Easy)
// https://leetcode.com/problems/ransom-note/
// Run: go run 0383_ransom_note.go lc.go
package main

import (
	"fmt"
	"os"
)

func canConstruct(ransomNote string, magazine string) bool {
	var have [26]int
	for _, c := range magazine {
		have[c-'a']++
	}
	for _, c := range ransomNote {
		if have[c-'a']--; have[c-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {
	assertEq(canConstruct("a", "b"), false)
	assertEq(canConstruct("aa", "ab"), false)
	assertEq(canConstruct("aa", "aab"), true)
	assertEq(canConstruct("", "x"), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
