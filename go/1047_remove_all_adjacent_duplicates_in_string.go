//go:build ignore

// 1047. Remove All Adjacent Duplicates In String (Easy)
// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
// Run: go run 1047_remove_all_adjacent_duplicates_in_string.go lc.go
package main

import (
	"fmt"
	"os"
)

func removeDuplicates(s string) string {
	st := []byte{}
	for i := 0; i < len(s); i++ {
		if len(st) > 0 && st[len(st)-1] == s[i] {
			st = st[:len(st)-1]
		} else {
			st = append(st, s[i])
		}
	}
	return string(st)
}

func main() {
	assertEq(removeDuplicates("abbaca"), "ca")
	assertEq(removeDuplicates("azxxzy"), "ay")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
