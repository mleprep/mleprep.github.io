//go:build ignore

// 877. Stone Game (Medium)
// https://leetcode.com/problems/stone-game/
// Run: go run 0877_stone_game.go lc.go
package main

import (
	"fmt"
	"os"
)

func stoneGame(piles []int) bool {
	n := len(piles)
	dp := append([]int(nil), piles...) // dp[i] = best (mine - theirs) on piles[i..i+length-1]
	for length := 2; length <= n; length++ {
		for i := 0; i+length <= n; i++ {
			j := i + length - 1
			dp[i] = max(piles[i]-dp[i+1], piles[j]-dp[i])
		}
	}
	return dp[0] > 0
}

func main() {
	assertEq(stoneGame([]int{5, 3, 4, 5}), true)
	assertEq(stoneGame([]int{3, 7, 2, 3}), true)
	assertEq(stoneGame([]int{1, 100, 2, 4}), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
