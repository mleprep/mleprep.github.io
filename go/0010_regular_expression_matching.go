//go:build ignore

// 10. Regular Expression Matching (Hard)
// https://leetcode.com/problems/regular-expression-matching/
// Run: go run 0010_regular_expression_matching.go lc.go
package main

import (
	"fmt"
	"os"
)

func isMatch(s string, p string) bool {
	memo := map[[2]int]bool{}
	var dp func(i, j int) bool // does s[i:] match p[j:]?
	dp = func(i, j int) bool {
		if j == len(p) {
			return i == len(s)
		}
		key := [2]int{i, j}
		if v, ok := memo[key]; ok {
			return v
		}
		first := i < len(s) && (p[j] == s[i] || p[j] == '.')
		var res bool
		if j+1 < len(p) && p[j+1] == '*' {
			res = dp(i, j+2) || (first && dp(i+1, j)) // skip "x*", or use it once more
		} else {
			res = first && dp(i+1, j+1)
		}
		memo[key] = res
		return res
	}
	return dp(0, 0)
}

func main() {
	assertEq(isMatch("aa", "a"), false)
	assertEq(isMatch("aa", "a*"), true)
	assertEq(isMatch("ab", ".*"), true)
	assertEq(isMatch("aab", "c*a*b"), true)
	assertEq(isMatch("mississippi", "mis*is*p*."), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
