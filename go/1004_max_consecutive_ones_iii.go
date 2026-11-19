//go:build ignore

// 1004. Max Consecutive Ones III (Medium)
// https://leetcode.com/problems/max-consecutive-ones-iii/
// Run: go run 1004_max_consecutive_ones_iii.go lc.go
package main

import (
	"fmt"
	"os"
)

func longestOnes(nums []int, k int) int {
	l, zeros, best := 0, 0, 0
	for r, x := range nums {
		if x == 0 {
			zeros++
		}
		for zeros > k {
			if nums[l] == 0 {
				zeros--
			}
			l++
		}
		best = max(best, r-l+1)
	}
	return best
}

func main() {
	assertEq(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2), 6)
	assertEq(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3), 10)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
