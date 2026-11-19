//go:build ignore

// 279. Perfect Squares (Medium)
// https://leetcode.com/problems/perfect-squares/
// Run: go run 0279_perfect_squares.go lc.go
package main

import (
	"fmt"
	"os"
)

func numSquares(n int) int {
	dp := make([]int, n+1) // dp[i] = fewest perfect squares summing to i
	for i := 1; i <= n; i++ {
		dp[i] = i // all ones
		for k := 1; k*k <= i; k++ {
			dp[i] = min(dp[i], dp[i-k*k]+1)
		}
	}
	return dp[n]
}

func main() {
	assertEq(numSquares(12), 3)
	assertEq(numSquares(13), 2)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
