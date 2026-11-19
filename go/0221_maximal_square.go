//go:build ignore

// 221. Maximal Square (Medium)
// https://leetcode.com/problems/maximal-square/
// Run: go run 0221_maximal_square.go lc.go
package main

import (
	"fmt"
	"os"
)

func maximalSquare(matrix [][]byte) int {
	n := len(matrix[0])
	dp, best := make([]int, n+1), 0 // dp[c] = side of the largest square ending at (r, c-1)
	for _, row := range matrix {
		diag := 0 // dp[r-1][c-1]
		for c := 1; c <= n; c++ {
			up := dp[c]
			if row[c-1] == '1' {
				dp[c] = 1 + min(dp[c], dp[c-1], diag)
			} else {
				dp[c] = 0
			}
			diag = up
			best = max(best, dp[c])
		}
	}
	return best * best
}

func main() {
	assertEq(maximalSquare(byteGrid("10100", "10111", "11111", "10010")), 4)
	assertEq(maximalSquare(byteGrid("0")), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
