//go:build ignore

// 219. Contains Duplicate II (Easy)
// https://leetcode.com/problems/contains-duplicate-ii/
// Run: go run 0219_contains_duplicate_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	window := map[int]bool{} // the last k values
	for i, x := range nums {
		if window[x] {
			return true
		}
		window[x] = true
		if len(window) > k {
			delete(window, nums[i-k])
		}
	}
	return false
}

func main() {
	assertEq(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3), true)
	assertEq(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1), true)
	assertEq(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2), false)
	assertEq(containsNearbyDuplicate([]int{1, 1}, 0), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
