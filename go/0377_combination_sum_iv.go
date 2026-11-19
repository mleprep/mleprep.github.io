//go:build ignore

// 377. Combination Sum IV (Medium)
// https://leetcode.com/problems/combination-sum-iv/
// Run: go run 0377_combination_sum_iv.go lc.go
package main

import (
	"fmt"
	"os"
)

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1) // dp[t] = ordered ways to sum to t
	dp[0] = 1
	for t := 1; t <= target; t++ {
		for _, x := range nums { // the last number of the sequence
			if x <= t {
				dp[t] += dp[t-x]
			}
		}
	}
	return dp[target]
}

func main() {
	assertEq(combinationSum4([]int{1, 2, 3}, 4), 7)
	assertEq(combinationSum4([]int{9}, 3), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
