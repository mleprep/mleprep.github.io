//go:build ignore

// 309. Best Time to Buy and Sell Stock with Cooldown (Medium)
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
// Run: go run 0309_best_time_to_buy_and_sell_stock_with_cooldown.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func maxProfit(prices []int) int {
	hold, sold, rest := math.MinInt/2, 0, 0 // holding a share / just sold / free to buy
	for _, p := range prices {
		hold, sold, rest = max(hold, rest-p), hold+p, max(rest, sold)
	}
	return max(sold, rest)
}

func main() {
	assertEq(maxProfit([]int{1, 2, 3, 0, 2}), 3)
	assertEq(maxProfit([]int{1}), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
