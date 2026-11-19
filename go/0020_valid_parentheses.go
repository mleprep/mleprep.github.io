//go:build ignore

// 20. Valid Parentheses (Easy)
// https://leetcode.com/problems/valid-parentheses/
// Run: go run 0020_valid_parentheses.go lc.go
package main

import (
	"fmt"
	"os"
)

func isValid(s string) bool {
	pair := map[byte]byte{')': '(', ']': '[', '}': '{'}
	st := []byte{}
	for i := 0; i < len(s); i++ {
		if open, ok := pair[s[i]]; ok {
			if len(st) == 0 || st[len(st)-1] != open {
				return false
			}
			st = st[:len(st)-1]
		} else {
			st = append(st, s[i])
		}
	}
	return len(st) == 0
}

func main() {
	assertEq(isValid("()[]{}"), true)
	assertEq(isValid("(]"), false)
	assertEq(isValid("([)]"), false)
	assertEq(isValid("{[]}"), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
