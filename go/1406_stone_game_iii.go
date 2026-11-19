//go:build ignore

// 1406. Stone Game III (Hard)
// https://leetcode.com/problems/stone-game-iii/
// Run: go run 1406_stone_game_iii.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	dp := make([]int, n+1) // dp[i] = best (mine - theirs) when it is my turn at i
	for i := n - 1; i >= 0; i-- {
		take, best := 0, math.MinInt
		for k := 0; k < 3 && i+k < n; k++ {
			take += stoneValue[i+k]
			best = max(best, take-dp[i+k+1])
		}
		dp[i] = best
	}
	switch {
	case dp[0] > 0:
		return "Alice"
	case dp[0] < 0:
		return "Bob"
	}
	return "Tie"
}

func main() {
	assertEq(stoneGameIII([]int{1, 2, 3, 7}), "Bob")
	assertEq(stoneGameIII([]int{1, 2, 3, -9}), "Alice")
	assertEq(stoneGameIII([]int{1, 2, 3, 6}), "Tie")
	assertEq(stoneGameIII([]int{-1, -2, -3}), "Tie")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
