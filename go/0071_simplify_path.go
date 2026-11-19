//go:build ignore

// 71. Simplify Path (Medium)
// https://leetcode.com/problems/simplify-path/
// Run: go run 0071_simplify_path.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func simplifyPath(path string) string {
	st := []string{}
	for _, part := range strings.Split(path, "/") {
		switch part {
		case "..":
			if len(st) > 0 {
				st = st[:len(st)-1]
			}
		case "", ".":
		default:
			st = append(st, part)
		}
	}
	return "/" + strings.Join(st, "/")
}

func main() {
	assertEq(simplifyPath("/home/"), "/home")
	assertEq(simplifyPath("/../"), "/")
	assertEq(simplifyPath("/home//foo/"), "/home/foo")
	assertEq(simplifyPath("/a/./b/../../c/"), "/c")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
