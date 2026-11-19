//go:build ignore

// 205. Isomorphic Strings (Easy)
// https://leetcode.com/problems/isomorphic-strings/
// Run: go run 0205_isomorphic_strings.go lc.go
package main

import (
	"fmt"
	"os"
)

func isIsomorphic(s string, t string) bool {
	var st, ts [256]byte // 0 = not mapped yet
	for i := 0; i < len(s); i++ {
		a, b := s[i], t[i]
		if st[a] == 0 && ts[b] == 0 {
			st[a], ts[b] = b, a
		} else if st[a] != b || ts[b] != a {
			return false
		}
	}
	return true
}

func main() {
	assertEq(isIsomorphic("egg", "add"), true)
	assertEq(isIsomorphic("foo", "bar"), false)
	assertEq(isIsomorphic("badc", "baba"), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
