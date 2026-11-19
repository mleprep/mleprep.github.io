//go:build ignore

// 704. Binary Search (Easy)
// https://leetcode.com/problems/binary-search/
// Run: go run 0704_binary_search.go lc.go
package main

import (
	"fmt"
	"os"
)

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return -1
}

func main() {
	assertEq(search([]int{-1, 0, 3, 5, 9, 12}, 9), 4)
	assertEq(search([]int{-1, 0, 3, 5, 9, 12}, 2), -1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
