//go:build ignore

// 301. Remove Invalid Parentheses (Hard)
// https://leetcode.com/problems/remove-invalid-parentheses/
// Run: go run 0301_remove_invalid_parentheses.go lc.go
package main

import (
	"fmt"
	"os"
)

func removeInvalidParentheses(s string) []string {
	valid := func(t string) bool {
		bal := 0
		for i := 0; i < len(t); i++ {
			if t[i] == '(' {
				bal++
			} else if t[i] == ')' {
				if bal--; bal < 0 {
					return false
				}
			}
		}
		return bal == 0
	}
	level := map[string]bool{s: true}
	for { // BFS by number of removals: the first level with a valid string wins
		good := []string{}
		for t := range level {
			if valid(t) {
				good = append(good, t)
			}
		}
		if len(good) > 0 {
			return good
		}
		next := map[string]bool{}
		for t := range level {
			for i := 0; i < len(t); i++ {
				if t[i] == '(' || t[i] == ')' {
					next[t[:i]+t[i+1:]] = true
				}
			}
		}
		level = next
	}
}

func main() {
	assertUnordered(removeInvalidParentheses("()())()"), []string{"(())()", "()()()"})
	assertUnordered(removeInvalidParentheses("(a)())()"), []string{"(a())()", "(a)()()"})
	assertEq(removeInvalidParentheses(")("), []string{""})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
