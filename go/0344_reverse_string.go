//go:build ignore

// 344. Reverse String (Easy)
// https://leetcode.com/problems/reverse-string/
// Run: go run 0344_reverse_string.go lc.go
package main

import (
	"fmt"
	"os"
)

func reverseString(s []byte) {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
}

func main() {
	a := []byte("hello")
	reverseString(a)
	assertEq(string(a), "olleh")
	b := []byte("Hannah")
	reverseString(b)
	assertEq(string(b), "hannaH")
	c := []byte("x")
	reverseString(c)
	assertEq(string(c), "x")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
