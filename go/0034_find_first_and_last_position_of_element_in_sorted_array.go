//go:build ignore

// 34. Find First and Last Position of Element in Sorted Array (Medium)
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// Run: go run 0034_find_first_and_last_position_of_element_in_sorted_array.go lc.go
package main

import (
	"fmt"
	"os"
)

func searchRange(nums []int, target int) []int {
	lower := func(t int) int { // first index with nums[i] >= t
		lo, hi := 0, len(nums)
		for lo < hi {
			mid := lo + (hi-lo)/2
			if nums[mid] >= t {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		return lo
	}
	a := lower(target)
	if a == len(nums) || nums[a] != target {
		return []int{-1, -1}
	}
	return []int{a, lower(target+1) - 1}
}

func main() {
	assertEq(searchRange([]int{5, 7, 7, 8, 8, 10}, 8), []int{3, 4})
	assertEq(searchRange([]int{5, 7, 7, 8, 8, 10}, 6), []int{-1, -1})
	assertEq(searchRange([]int{}, 0), []int{-1, -1})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
