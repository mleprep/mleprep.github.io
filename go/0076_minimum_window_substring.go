//go:build ignore

// 76. Minimum Window Substring (Hard)
// https://leetcode.com/problems/minimum-window-substring/
// Run: go run 0076_minimum_window_substring.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func minWindow(s string, t string) string {
	var need [128]int
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	missing, l, start, best := len(t), 0, 0, math.MaxInt
	for r := 0; r < len(s); r++ {
		if need[s[r]] > 0 {
			missing--
		}
		need[s[r]]--
		for missing == 0 { // window covers t: shrink it
			if r-l+1 < best {
				start, best = l, r-l+1
			}
			need[s[l]]++
			if need[s[l]] > 0 {
				missing++
			}
			l++
		}
	}
	if best == math.MaxInt {
		return ""
	}
	return s[start : start+best]
}

func main() {
	assertEq(minWindow("ADOBECODEBANC", "ABC"), "BANC")
	assertEq(minWindow("a", "aa"), "")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
