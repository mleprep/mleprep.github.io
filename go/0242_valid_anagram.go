//go:build ignore

// 242. Valid Anagram (Easy)
// https://leetcode.com/problems/valid-anagram/
// Run: go run 0242_valid_anagram.go lc.go
package main

import (
	"fmt"
	"os"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var cnt [26]int
	for i := range s {
		cnt[s[i]-'a']++
		cnt[t[i]-'a']--
	}
	return cnt == [26]int{} // arrays compare by value
}

func main() {
	assertEq(isAnagram("anagram", "nagaram"), true)
	assertEq(isAnagram("rat", "car"), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
