//go:build ignore

// 128. Longest Consecutive Sequence (Medium)
// https://leetcode.com/problems/longest-consecutive-sequence/
// Run: go run 0128_longest_consecutive_sequence.go lc.go
package main

import (
	"fmt"
	"os"
)

func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, x := range nums {
		set[x] = true
	}
	best := 0
	for x := range set {
		if !set[x-1] { // x starts a run
			y := x
			for set[y+1] {
				y++
			}
			best = max(best, y-x+1)
		}
	}
	return best
}

func main() {
	assertEq(longestConsecutive([]int{100, 4, 200, 1, 3, 2}), 4)
	assertEq(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}), 9)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
