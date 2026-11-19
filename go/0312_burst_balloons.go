//go:build ignore

// 312. Burst Balloons (Hard)
// https://leetcode.com/problems/burst-balloons/
// Run: go run 0312_burst_balloons.go lc.go
package main

import (
	"fmt"
	"os"
)

func maxCoins(nums []int) int {
	a := append(append([]int{1}, nums...), 1)
	n := len(a)
	dp := make([][]int, n) // dp[l][r] = best for the balloons strictly between l and r
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for length := 2; length < n; length++ {
		for l := 0; l+length < n; l++ {
			r := l + length
			for k := l + 1; k < r; k++ { // k is the last balloon burst between l and r
				dp[l][r] = max(dp[l][r], dp[l][k]+a[l]*a[k]*a[r]+dp[k][r])
			}
		}
	}
	return dp[0][n-1]
}

func main() {
	assertEq(maxCoins([]int{3, 1, 5, 8}), 167)
	assertEq(maxCoins([]int{1, 5}), 10)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
