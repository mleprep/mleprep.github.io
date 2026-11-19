//go:build ignore

// 983. Minimum Cost For Tickets (Medium)
// https://leetcode.com/problems/minimum-cost-for-tickets/
// Run: go run 0983_minimum_cost_for_tickets.go lc.go
package main

import (
	"fmt"
	"os"
)

func mincostTickets(days []int, costs []int) int {
	last := days[len(days)-1]
	travel := make([]bool, last+1)
	for _, d := range days {
		travel[d] = true
	}
	dp := make([]int, last+1) // dp[d] = cheapest cover of all travel days up to d
	for d := 1; d <= last; d++ {
		if !travel[d] {
			dp[d] = dp[d-1]
			continue
		}
		dp[d] = min(dp[d-1]+costs[0], dp[max(0, d-7)]+costs[1], dp[max(0, d-30)]+costs[2])
	}
	return dp[last]
}

func main() {
	assertEq(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}), 11)
	assertEq(mincostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}), 17)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
