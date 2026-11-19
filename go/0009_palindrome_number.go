//go:build ignore

// 9. Palindrome Number (Easy)
// https://leetcode.com/problems/palindrome-number/
// Run: go run 0009_palindrome_number.go lc.go
package main

import (
	"fmt"
	"os"
)

func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}
	rev := 0
	for x > rev { // reverse only the second half
		rev = rev*10 + x%10
		x /= 10
	}
	return x == rev || x == rev/10 // odd length: drop the middle digit
}

func main() {
	assertEq(isPalindrome(121), true)
	assertEq(isPalindrome(-121), false)
	assertEq(isPalindrome(10), false)
	assertEq(isPalindrome(1221), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
