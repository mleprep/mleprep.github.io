//go:build ignore

// 1249. Minimum Remove to Make Valid Parentheses (Medium)
// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/
// Run: go run 1249_minimum_remove_to_make_valid_parentheses.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func minRemoveToMakeValid(s string) string {
	st, bad := []int{}, map[int]bool{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			st = append(st, i)
		} else if s[i] == ')' {
			if len(st) > 0 {
				st = st[:len(st)-1]
			} else {
				bad[i] = true
			}
		}
	}
	for _, i := range st { // unmatched "("
		bad[i] = true
	}
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		if !bad[i] {
			sb.WriteByte(s[i])
		}
	}
	return sb.String()
}

func main() {
	assertEq(minRemoveToMakeValid("lee(t(c)o)de)"), "lee(t(c)o)de")
	assertEq(minRemoveToMakeValid("a)b(c)d"), "ab(c)d")
	assertEq(minRemoveToMakeValid("))(("), "")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
