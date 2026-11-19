//go:build ignore

// 518. Coin Change II (Medium)
// https://leetcode.com/problems/coin-change-ii/
// Run: go run 0518_coin_change_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func change(amount int, coins []int) int {
	dp := make([]int, amount+1) // dp[a] = combinations summing to a
	dp[0] = 1
	for _, c := range coins { // coins in the outer loop: each combination counted once
		for a := c; a <= amount; a++ {
			dp[a] += dp[a-c]
		}
	}
	return dp[amount]
}

func main() {
	assertEq(change(5, []int{1, 2, 5}), 4)
	assertEq(change(3, []int{2}), 0)
	assertEq(change(10, []int{10}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
