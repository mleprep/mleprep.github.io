//go:build ignore

// 300. Longest Increasing Subsequence (Medium)
// https://leetcode.com/problems/longest-increasing-subsequence/
// Run: go run 0300_longest_increasing_subsequence.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func lengthOfLIS(nums []int) int {
	tails := []int{} // tails[k] = smallest tail of an increasing subsequence of length k+1
	for _, x := range nums {
		i := sort.SearchInts(tails, x)
		if i == len(tails) {
			tails = append(tails, x)
		} else {
			tails[i] = x
		}
	}
	return len(tails)
}

func main() {
	assertEq(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}), 4)
	assertEq(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}), 4)
	assertEq(lengthOfLIS([]int{7, 7, 7, 7}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
