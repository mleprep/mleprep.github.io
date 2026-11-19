//go:build ignore

// 33. Search in Rotated Sorted Array (Medium)
// https://leetcode.com/problems/search-in-rotated-sorted-array/
// Run: go run 0033_search_in_rotated_sorted_array.go lc.go
package main

import (
	"fmt"
	"os"
)

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		if nums[lo] <= nums[mid] { // left half sorted
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else { // right half sorted
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

func main() {
	assertEq(search([]int{4, 5, 6, 7, 0, 1, 2}, 0), 4)
	assertEq(search([]int{4, 5, 6, 7, 0, 1, 2}, 3), -1)
	assertEq(search([]int{1}, 0), -1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
