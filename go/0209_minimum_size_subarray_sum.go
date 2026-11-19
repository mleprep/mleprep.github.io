//go:build ignore

// 209. Minimum Size Subarray Sum (Medium)
// https://leetcode.com/problems/minimum-size-subarray-sum/
// Run: go run 0209_minimum_size_subarray_sum.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func minSubArrayLen(target int, nums []int) int {
	l, s, best := 0, 0, math.MaxInt
	for r, x := range nums {
		s += x
		for s >= target {
			best = min(best, r-l+1)
			s -= nums[l]
			l++
		}
	}
	if best == math.MaxInt {
		return 0
	}
	return best
}

func main() {
	assertEq(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}), 2)
	assertEq(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
