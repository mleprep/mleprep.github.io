//go:build ignore

// 322. Coin Change (Medium)
// https://leetcode.com/problems/coin-change/
// Run: go run 0322_coin_change.go lc.go
package main

import (
	"fmt"
	"os"
)

func coinChange(coins []int, amount int) int {
	inf := amount + 1
	dp := make([]int, amount+1) // dp[a] = fewest coins summing to a
	for a := 1; a <= amount; a++ {
		dp[a] = inf
		for _, c := range coins {
			if c <= a {
				dp[a] = min(dp[a], dp[a-c]+1)
			}
		}
	}
	if dp[amount] == inf {
		return -1
	}
	return dp[amount]
}

func main() {
	assertEq(coinChange([]int{1, 2, 5}, 11), 3)
	assertEq(coinChange([]int{2}, 3), -1)
	assertEq(coinChange([]int{1}, 0), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
