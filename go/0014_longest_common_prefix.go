//go:build ignore

// 14. Longest Common Prefix (Easy)
// https://leetcode.com/problems/longest-common-prefix/
// Run: go run 0014_longest_common_prefix.go lc.go
package main

import (
	"fmt"
	"os"
)

func longestCommonPrefix(strs []string) string {
	first := strs[0]
	for i := 0; i < len(first); i++ {
		for _, s := range strs[1:] {
			if i == len(s) || s[i] != first[i] {
				return first[:i]
			}
		}
	}
	return first
}

func main() {
	assertEq(longestCommonPrefix([]string{"flower", "flow", "flight"}), "fl")
	assertEq(longestCommonPrefix([]string{"dog", "racecar", "car"}), "")
	assertEq(longestCommonPrefix([]string{"ab", "abc"}), "ab")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
