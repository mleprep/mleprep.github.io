//go:build ignore

// 1493. Longest Subarray of 1's After Deleting One Element (Medium)
// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/
// Run: go run 1493_longest_subarray_of_1s_after_deleting_one_element.go lc.go
package main

import (
	"fmt"
	"os"
)

func longestSubarray(nums []int) int {
	lo, zeros, best := 0, 0, 0
	for hi, x := range nums {
		if x == 0 {
			zeros++
		}
		for zeros > 1 {
			if nums[lo] == 0 {
				zeros--
			}
			lo++
		}
		best = max(best, hi-lo) // (hi - lo + 1) - 1 deleted element
	}
	return best
}

func main() {
	assertEq(longestSubarray([]int{1, 1, 0, 1}), 3)
	assertEq(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}), 5)
	assertEq(longestSubarray([]int{1, 1, 1}), 2)
	assertEq(longestSubarray([]int{0, 0, 0}), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
