//go:build ignore

// 659. Split Array into Consecutive Subsequences (Medium)
// https://leetcode.com/problems/split-array-into-consecutive-subsequences/
// Run: go run 0659_split_array_into_consecutive_subsequences.go lc.go
package main

import (
	"fmt"
	"os"
)

func isPossible(nums []int) bool {
	left, ends := map[int]int{}, map[int]int{} // unused counts; subsequences ending at x
	for _, x := range nums {
		left[x]++
	}
	for _, x := range nums {
		if left[x] == 0 {
			continue
		}
		left[x]--
		if ends[x-1] > 0 { // extend an existing subsequence
			ends[x-1]--
			ends[x]++
		} else if left[x+1] > 0 && left[x+2] > 0 { // or start a new one of length 3
			left[x+1]--
			left[x+2]--
			ends[x+2]++
		} else {
			return false
		}
	}
	return true
}

func main() {
	assertEq(isPossible([]int{1, 2, 3, 3, 4, 5}), true)
	assertEq(isPossible([]int{1, 2, 3, 3, 4, 4, 5, 5}), true)
	assertEq(isPossible([]int{1, 2, 3, 4, 4, 5}), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
