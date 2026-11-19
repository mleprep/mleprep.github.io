//go:build ignore

// 121. Best Time to Buy and Sell Stock (Easy)
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// Run: go run 0121_best_time_to_buy_and_sell_stock.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func maxProfit(prices []int) int {
	lo, best := math.MaxInt, 0
	for _, p := range prices {
		lo = min(lo, p) // cheapest buy so far
		best = max(best, p-lo)
	}
	return best
}

func main() {
	assertEq(maxProfit([]int{7, 1, 5, 3, 6, 4}), 5)
	assertEq(maxProfit([]int{7, 6, 4, 3, 1}), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
