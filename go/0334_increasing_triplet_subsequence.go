//go:build ignore

// 334. Increasing Triplet Subsequence (Medium)
// https://leetcode.com/problems/increasing-triplet-subsequence/
// Run: go run 0334_increasing_triplet_subsequence.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt, math.MaxInt // smallest end of increasing runs of length 1 and 2
	for _, x := range nums {
		if x <= first {
			first = x
		} else if x <= second {
			second = x
		} else {
			return true
		}
	}
	return false
}

func main() {
	assertEq(increasingTriplet([]int{1, 2, 3, 4, 5}), true)
	assertEq(increasingTriplet([]int{5, 4, 3, 2, 1}), false)
	assertEq(increasingTriplet([]int{2, 1, 5, 0, 4, 6}), true)
	assertEq(increasingTriplet([]int{20, 100, 10, 12, 5, 13}), true)
	assertEq(increasingTriplet([]int{1, 1, 1}), false)
	assertEq(increasingTriplet([]int{1, 5, 0, 4, 1, 3}), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
