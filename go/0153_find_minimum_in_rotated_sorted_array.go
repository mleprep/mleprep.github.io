//go:build ignore

// 153. Find Minimum in Rotated Sorted Array (Medium)
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
// Run: go run 0153_find_minimum_in_rotated_sorted_array.go lc.go
package main

import (
	"fmt"
	"os"
)

func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[hi] { // the drop is to the right of mid
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return nums[lo]
}

func main() {
	assertEq(findMin([]int{3, 4, 5, 1, 2}), 1)
	assertEq(findMin([]int{4, 5, 6, 7, 0, 1, 2}), 0)
	assertEq(findMin([]int{11, 13, 15, 17}), 11)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
