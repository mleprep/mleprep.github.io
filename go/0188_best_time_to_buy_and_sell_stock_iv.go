//go:build ignore

// 188. Best Time to Buy and Sell Stock IV (Hard)
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
// Run: go run 0188_best_time_to_buy_and_sell_stock_iv.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func maxProfit(k int, prices []int) int {
	if k >= len(prices)/2 { // effectively unlimited: take every rise
		total := 0
		for i := 1; i < len(prices); i++ {
			total += max(0, prices[i]-prices[i-1])
		}
		return total
	}
	buy, sell := make([]int, k+1), make([]int, k+1) // after the j-th buy / j-th sell
	for j := range buy {
		buy[j] = math.MinInt / 2
	}
	for _, p := range prices {
		for j := 1; j <= k; j++ {
			buy[j] = max(buy[j], sell[j-1]-p)
			sell[j] = max(sell[j], buy[j]+p)
		}
	}
	return sell[k]
}

func main() {
	assertEq(maxProfit(2, []int{2, 4, 1}), 2)
	assertEq(maxProfit(2, []int{3, 2, 6, 5, 0, 3}), 7)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
