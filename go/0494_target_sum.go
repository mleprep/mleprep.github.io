//go:build ignore

// 494. Target Sum (Medium)
// https://leetcode.com/problems/target-sum/
// Run: go run 0494_target_sum.go lc.go
package main

import (
	"fmt"
	"os"
)

func findTargetSumWays(nums []int, target int) int {
	total := 0
	for _, x := range nums {
		total += x
	}
	if target > total || -target > total || (total+target)%2 != 0 {
		return 0
	}
	P := (total + target) / 2 // the '+' numbers must sum to P
	dp := make([]int, P+1)
	dp[0] = 1
	for _, x := range nums {
		for s := P; s >= x; s-- {
			dp[s] += dp[s-x]
		}
	}
	return dp[P]
}

func main() {
	assertEq(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3), 5)
	assertEq(findTargetSumWays([]int{1}, 1), 1)
	assertEq(findTargetSumWays([]int{0, 0, 1}, 1), 4)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
