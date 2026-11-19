//go:build ignore

// 724. Find Pivot Index (Easy)
// https://leetcode.com/problems/find-pivot-index/
// Run: go run 0724_find_pivot_index.go lc.go
package main

import (
	"fmt"
	"os"
)

func pivotIndex(nums []int) int {
	total, left := 0, 0
	for _, x := range nums {
		total += x
	}
	for i, x := range nums {
		if left == total-left-x { // left sum equals right sum
			return i
		}
		left += x
	}
	return -1
}

func main() {
	assertEq(pivotIndex([]int{1, 7, 3, 6, 5, 6}), 3)
	assertEq(pivotIndex([]int{1, 2, 3}), -1)
	assertEq(pivotIndex([]int{2, 1, -1}), 0)
	assertEq(pivotIndex([]int{0}), 0)
	assertEq(pivotIndex([]int{-1, -1, 0, 1, 1, 0}), 5)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
