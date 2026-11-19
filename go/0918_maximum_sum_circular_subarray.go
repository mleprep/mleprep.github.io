//go:build ignore

// 918. Maximum Sum Circular Subarray (Medium)
// https://leetcode.com/problems/maximum-sum-circular-subarray/
// Run: go run 0918_maximum_sum_circular_subarray.go lc.go
package main

import (
	"fmt"
	"os"
)

func maxSubarraySumCircular(nums []int) int {
	curMax, curMin, total := 0, 0, 0
	bestMax, bestMin := nums[0], nums[0]
	for _, x := range nums {
		curMax = max(x, curMax+x)
		bestMax = max(bestMax, curMax)
		curMin = min(x, curMin+x)
		bestMin = min(bestMin, curMin)
		total += x
	}
	if bestMax < 0 {
		return bestMax // all negative: the wrap-around case would be empty
	}
	return max(bestMax, total-bestMin) // a wrapping subarray = total minus a middle one
}

func main() {
	assertEq(maxSubarraySumCircular([]int{1, -2, 3, -2}), 3)
	assertEq(maxSubarraySumCircular([]int{5, -3, 5}), 10)
	assertEq(maxSubarraySumCircular([]int{-3, -2, -3}), -2)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
