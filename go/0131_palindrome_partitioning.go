//go:build ignore

// 131. Palindrome Partitioning (Medium)
// https://leetcode.com/problems/palindrome-partitioning/
// Run: go run 0131_palindrome_partitioning.go lc.go
package main

import (
	"fmt"
	"os"
)

func partition(s string) [][]string {
	isPal := func(i, j int) bool {
		for ; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}
	res, path := [][]string{}, []string{}
	var bt func(i int)
	bt = func(i int) {
		if i == len(s) {
			res = append(res, append([]string{}, path...))
			return
		}
		for j := i; j < len(s); j++ {
			if isPal(i, j) {
				path = append(path, s[i:j+1])
				bt(j + 1)
				path = path[:len(path)-1]
			}
		}
	}
	bt(0)
	return res
}

func main() {
	assertUnordered(partition("aab"), [][]string{{"a", "a", "b"}, {"aa", "b"}})
	assertEq(partition("a"), [][]string{{"a"}})
	assertUnordered(partition("aba"), [][]string{{"a", "b", "a"}, {"aba"}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
