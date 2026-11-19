//go:build ignore

// 714. Best Time to Buy and Sell Stock with Transaction Fee (Medium)
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
// Run: go run 0714_best_time_to_buy_and_sell_stock_with_transaction_fee.go lc.go
package main

import (
	"fmt"
	"os"
)

func maxProfit(prices []int, fee int) int {
	cash, hold := 0, -prices[0] // best profit without / with a share in hand
	for _, p := range prices[1:] {
		cash = max(cash, hold+p-fee)
		hold = max(hold, cash-p)
	}
	return cash
}

func main() {
	assertEq(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2), 8)
	assertEq(maxProfit([]int{1, 3, 7, 5, 10, 3}, 3), 6)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
