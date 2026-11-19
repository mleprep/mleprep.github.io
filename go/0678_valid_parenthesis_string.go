//go:build ignore

// 678. Valid Parenthesis String (Medium)
// https://leetcode.com/problems/valid-parenthesis-string/
// Run: go run 0678_valid_parenthesis_string.go lc.go
package main

import (
	"fmt"
	"os"
)

func checkValidString(s string) bool {
	lo, hi := 0, 0 // range of possible open-bracket counts
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			lo, hi = lo+1, hi+1
		case ')':
			lo, hi = lo-1, hi-1
		default: // '*' is ')', '' or '('
			lo, hi = lo-1, hi+1
		}
		if hi < 0 {
			return false
		}
		lo = max(lo, 0)
	}
	return lo == 0
}

func main() {
	assertEq(checkValidString("()"), true)
	assertEq(checkValidString("(*)"), true)
	assertEq(checkValidString("(*))"), true)
	assertEq(checkValidString("(((*)"), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
