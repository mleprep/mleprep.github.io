//go:build ignore

// 123. Best Time to Buy and Sell Stock III (Hard)
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
// Run: go run 0123_best_time_to_buy_and_sell_stock_iii.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func maxProfit(prices []int) int {
	buy1, buy2 := math.MinInt/2, math.MinInt/2
	sell1, sell2 := 0, 0
	for _, p := range prices {
		buy1 = max(buy1, -p)
		sell1 = max(sell1, buy1+p)
		buy2 = max(buy2, sell1-p) // the second buy is paid from the first profit
		sell2 = max(sell2, buy2+p)
	}
	return sell2
}

func main() {
	assertEq(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}), 6)
	assertEq(maxProfit([]int{1, 2, 3, 4, 5}), 4)
	assertEq(maxProfit([]int{7, 6, 4, 3, 1}), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
