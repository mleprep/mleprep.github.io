//go:build ignore

// 32. Longest Valid Parentheses (Hard)
// https://leetcode.com/problems/longest-valid-parentheses/
// Run: go run 0032_longest_valid_parentheses.go lc.go
package main

import (
	"fmt"
	"os"
)

func longestValidParentheses(s string) int {
	st, best := []int{-1}, 0 // stack bottom = index just before the current valid run
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			st = append(st, i)
			continue
		}
		st = st[:len(st)-1]
		if len(st) == 0 {
			st = append(st, i) // unmatched ')': new base
		} else {
			best = max(best, i-st[len(st)-1])
		}
	}
	return best
}

func main() {
	assertEq(longestValidParentheses("(()"), 2)
	assertEq(longestValidParentheses(")()())"), 4)
	assertEq(longestValidParentheses(""), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
