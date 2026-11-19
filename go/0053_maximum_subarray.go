//go:build ignore

// 53. Maximum Subarray (Medium)
// https://leetcode.com/problems/maximum-subarray/
// Run: go run 0053_maximum_subarray.go lc.go
package main

import (
	"fmt"
	"os"
)

func maxSubArray(nums []int) int {
	best, cur := nums[0], nums[0] // cur = best sum ending here (Kadane)
	for _, x := range nums[1:] {
		cur = max(x, cur+x)
		best = max(best, cur)
	}
	return best
}

func main() {
	assertEq(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6)
	assertEq(maxSubArray([]int{-1}), -1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
