//go:build ignore

// 474. Ones and Zeroes (Medium)
// https://leetcode.com/problems/ones-and-zeroes/
// Run: go run 0474_ones_and_zeroes.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1) // dp[i][j] = most strings using at most i zeros and j ones
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		z := strings.Count(s, "0")
		o := len(s) - z
		for i := m; i >= z; i-- { // 0/1 knapsack in two dimensions, iterated downwards
			for j := n; j >= o; j-- {
				dp[i][j] = max(dp[i][j], dp[i-z][j-o]+1)
			}
		}
	}
	return dp[m][n]
}

func main() {
	assertEq(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3), 4)
	assertEq(findMaxForm([]string{"10", "0", "1"}, 1, 1), 2)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
