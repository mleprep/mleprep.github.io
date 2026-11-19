//go:build ignore

// 680. Valid Palindrome II (Easy)
// https://leetcode.com/problems/valid-palindrome-ii/
// Run: go run 0680_valid_palindrome_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func validPalindrome(s string) bool {
	pal := func(l, r int) bool {
		for ; l < r; l, r = l+1, r-1 {
			if s[l] != s[r] {
				return false
			}
		}
		return true
	}
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] { // delete one of the two
			return pal(l+1, r) || pal(l, r-1)
		}
	}
	return true
}

func main() {
	assertEq(validPalindrome("aba"), true)
	assertEq(validPalindrome("abca"), true)
	assertEq(validPalindrome("abc"), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
