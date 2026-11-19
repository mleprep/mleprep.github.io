//go:build ignore

// 81. Search in Rotated Sorted Array II (Medium)
// https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
// Run: go run 0081_search_in_rotated_sorted_array_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func search(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return true
		}
		switch {
		case nums[lo] == nums[mid] && nums[mid] == nums[hi]: // can't tell which half is sorted
			lo, hi = lo+1, hi-1
		case nums[lo] <= nums[mid]:
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		default:
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return false
}

func main() {
	assertEq(search([]int{2, 5, 6, 0, 0, 1, 2}, 0), true)
	assertEq(search([]int{2, 5, 6, 0, 0, 1, 2}, 3), false)
	assertEq(search([]int{1, 0, 1, 1, 1}, 0), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
