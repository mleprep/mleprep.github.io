//go:build ignore

// 343. Integer Break (Medium)
// https://leetcode.com/problems/integer-break/
// Run: go run 0343_integer_break.go lc.go
package main

import (
	"fmt"
	"os"
)

func integerBreak(n int) int {
	dp := make([]int, n+1) // dp[i] = best product splitting i into at least two parts
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ { // first part i-j; the rest j either kept whole or split further
			dp[i] = max(dp[i], max(j, dp[j])*(i-j))
		}
	}
	return dp[n]
}

func main() {
	assertEq(integerBreak(2), 1)
	assertEq(integerBreak(10), 36)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
