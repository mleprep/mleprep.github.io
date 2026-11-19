//go:build ignore

// 844. Backspace String Compare (Easy)
// https://leetcode.com/problems/backspace-string-compare/
// Run: go run 0844_backspace_string_compare.go lc.go
package main

import (
	"fmt"
	"os"
)

func backspaceCompare(s string, t string) bool {
	// next returns the index of the next surviving character at or before i (-1 if none)
	next := func(x string, i int) int {
		skip := 0
		for ; i >= 0; i-- {
			if x[i] == '#' {
				skip++
			} else if skip > 0 {
				skip--
			} else {
				break
			}
		}
		return i
	}
	i, j := len(s)-1, len(t)-1
	for {
		i, j = next(s, i), next(t, j)
		if i < 0 || j < 0 {
			return i < 0 && j < 0
		}
		if s[i] != t[j] {
			return false
		}
		i, j = i-1, j-1
	}
}

func main() {
	assertEq(backspaceCompare("ab#c", "ad#c"), true)
	assertEq(backspaceCompare("ab##", "c#d#"), true)
	assertEq(backspaceCompare("a#c", "b"), false)
	assertEq(backspaceCompare("a##c", "#a#c"), true)
	assertEq(backspaceCompare("bxj##tw", "bxo#j##tw"), true)
	assertEq(backspaceCompare("ab", "b"), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
