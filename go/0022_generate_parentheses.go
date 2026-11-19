//go:build ignore

// 22. Generate Parentheses (Medium)
// https://leetcode.com/problems/generate-parentheses/
// Run: go run 0022_generate_parentheses.go lc.go
package main

import (
	"fmt"
	"os"
)

func generateParenthesis(n int) []string {
	res := []string{}
	var bt func(cur string, open, closed int)
	bt = func(cur string, open, closed int) {
		if len(cur) == 2*n {
			res = append(res, cur)
			return
		}
		if open < n {
			bt(cur+"(", open+1, closed)
		}
		if closed < open { // only close what is open
			bt(cur+")", open, closed+1)
		}
	}
	bt("", 0, 0)
	return res
}

func main() {
	assertUnordered(generateParenthesis(3), []string{"((()))", "(()())", "(())()", "()(())", "()()()"})
	assertEq(generateParenthesis(1), []string{"()"})
	assertUnordered(generateParenthesis(4), []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
