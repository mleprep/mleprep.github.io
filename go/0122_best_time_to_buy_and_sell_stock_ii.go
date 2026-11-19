//go:build ignore

// 122. Best Time to Buy and Sell Stock II (Medium)
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
// Run: go run 0122_best_time_to_buy_and_sell_stock_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func maxProfit(prices []int) int {
	total := 0
	for i := 1; i < len(prices); i++ {
		total += max(0, prices[i]-prices[i-1]) // collect every rise
	}
	return total
}

func main() {
	assertEq(maxProfit([]int{7, 1, 5, 3, 6, 4}), 7)
	assertEq(maxProfit([]int{1, 2, 3, 4, 5}), 4)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
