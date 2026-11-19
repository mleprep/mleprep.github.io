//go:build ignore

// 643. Maximum Average Subarray I (Easy)
// https://leetcode.com/problems/maximum-average-subarray-i/
// Run: go run 0643_maximum_average_subarray_i.go lc.go
package main

import (
	"fmt"
	"os"
)

func findMaxAverage(nums []int, k int) float64 {
	s := 0
	for _, x := range nums[:k] {
		s += x
	}
	best := s
	for i := k; i < len(nums); i++ {
		s += nums[i] - nums[i-k]
		best = max(best, s)
	}
	return float64(best) / float64(k)
}

func main() {
	assertEq(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4), 12.75)
	assertEq(findMaxAverage([]int{5}, 1), 5.0)
	assertEq(findMaxAverage([]int{0, 4, 0, 3, 2}, 1), 4.0)
	assertEq(findMaxAverage([]int{-1, -2, -3}, 3), -2.0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
