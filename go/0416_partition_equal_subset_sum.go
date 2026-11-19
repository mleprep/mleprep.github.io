//go:build ignore

// 416. Partition Equal Subset Sum (Medium)
// https://leetcode.com/problems/partition-equal-subset-sum/
// Run: go run 0416_partition_equal_subset_sum.go lc.go
package main

import (
	"fmt"
	"os"
)

func canPartition(nums []int) bool {
	total := 0
	for _, x := range nums {
		total += x
	}
	if total%2 == 1 {
		return false
	}
	target := total / 2
	dp := make([]bool, target+1) // dp[s]: some subset sums to s
	dp[0] = true
	for _, x := range nums {
		for s := target; s >= x; s-- { // downwards so each number is used once
			dp[s] = dp[s] || dp[s-x]
		}
	}
	return dp[target]
}

func main() {
	assertEq(canPartition([]int{1, 5, 11, 5}), true)
	assertEq(canPartition([]int{1, 2, 3, 5}), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
