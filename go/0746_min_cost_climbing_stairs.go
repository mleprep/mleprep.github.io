//go:build ignore

// 746. Min Cost Climbing Stairs (Easy)
// https://leetcode.com/problems/min-cost-climbing-stairs/
// Run: go run 0746_min_cost_climbing_stairs.go lc.go
package main

import (
	"fmt"
	"os"
)

func minCostClimbingStairs(cost []int) int {
	a, b := 0, 0 // dp[i-2], dp[i-1]: cheapest way to stand on step i
	for i := 2; i <= len(cost); i++ {
		a, b = b, min(b+cost[i-1], a+cost[i-2])
	}
	return b
}

func main() {
	assertEq(minCostClimbingStairs([]int{10, 15, 20}), 15)
	assertEq(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}), 6)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
