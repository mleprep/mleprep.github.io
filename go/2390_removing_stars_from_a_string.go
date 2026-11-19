//go:build ignore

// 2390. Removing Stars From a String (Medium)
// https://leetcode.com/problems/removing-stars-from-a-string/
// Run: go run 2390_removing_stars_from_a_string.go lc.go
package main

import (
	"fmt"
	"os"
)

func removeStars(s string) string {
	st := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			st = st[:len(st)-1]
		} else {
			st = append(st, s[i])
		}
	}
	return string(st)
}

func main() {
	assertEq(removeStars("leet**cod*e"), "lecoe")
	assertEq(removeStars("erase*****"), "")
	assertEq(removeStars("abc"), "abc")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
