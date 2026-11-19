//go:build ignore

// 35. Search Insert Position (Easy)
// https://leetcode.com/problems/search-insert-position/
// Run: go run 0035_search_insert_position.go lc.go
package main

import (
	"fmt"
	"os"
)

func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums) // first index with nums[i] >= target
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	assertEq(searchInsert([]int{1, 3, 5, 6}, 5), 2)
	assertEq(searchInsert([]int{1, 3, 5, 6}, 2), 1)
	assertEq(searchInsert([]int{1, 3, 5, 6}, 7), 4)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
