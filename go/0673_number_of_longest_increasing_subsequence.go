//go:build ignore

// 673. Number of Longest Increasing Subsequence (Medium)
// https://leetcode.com/problems/number-of-longest-increasing-subsequence/
// Run: go run 0673_number_of_longest_increasing_subsequence.go lc.go
package main

import (
	"fmt"
	"os"
)

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	length, cnt := make([]int, n), make([]int, n) // LIS length ending at i, and how many
	best := 0
	for i := 0; i < n; i++ {
		length[i], cnt[i] = 1, 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if length[j]+1 > length[i] {
					length[i], cnt[i] = length[j]+1, cnt[j]
				} else if length[j]+1 == length[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		best = max(best, length[i])
	}
	total := 0
	for i := range nums {
		if length[i] == best {
			total += cnt[i]
		}
	}
	return total
}

func main() {
	assertEq(findNumberOfLIS([]int{1, 3, 5, 4, 7}), 2)
	assertEq(findNumberOfLIS([]int{2, 2, 2, 2, 2}), 5)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
