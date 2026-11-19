//go:build ignore

// 125. Valid Palindrome (Easy)
// https://leetcode.com/problems/valid-palindrome/
// Run: go run 0125_valid_palindrome.go lc.go
package main

import (
	"fmt"
	"os"
)

func isPalindrome(s string) bool {
	alnum := func(c byte) bool { return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' }
	lower := func(c byte) byte {
		if c >= 'A' && c <= 'Z' {
			return c + 'a' - 'A'
		}
		return c
	}
	for l, r := 0, len(s)-1; l < r; {
		switch {
		case !alnum(s[l]):
			l++
		case !alnum(s[r]):
			r--
		case lower(s[l]) != lower(s[r]):
			return false
		default:
			l, r = l+1, r-1
		}
	}
	return true
}

func main() {
	assertEq(isPalindrome("A man, a plan, a canal: Panama"), true)
	assertEq(isPalindrome("race a car"), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
