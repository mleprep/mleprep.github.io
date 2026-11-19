//go:build ignore

// 1140. Stone Game II (Medium)
// https://leetcode.com/problems/stone-game-ii/
// Run: go run 1140_stone_game_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func stoneGameII(piles []int) int {
	n := len(piles)
	suffix := make([]int, n+1) // stones in piles[i:]
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + piles[i]
	}
	memo := map[[2]int]int{}
	var best func(i, m int) int // most stones the player to move gets from piles[i:]
	best = func(i, m int) int {
		if i+2*m >= n {
			return suffix[i] // take everything
		}
		key := [2]int{i, m}
		if v, ok := memo[key]; ok {
			return v
		}
		res := 0
		for x := 1; x <= 2*m; x++ {
			res = max(res, suffix[i]-best(i+x, max(m, x)))
		}
		memo[key] = res
		return res
	}
	return best(0, 1)
}

func main() {
	assertEq(stoneGameII([]int{2, 7, 9, 4, 4}), 10)
	assertEq(stoneGameII([]int{1, 2, 3, 4, 5, 100}), 104)
	assertEq(stoneGameII([]int{5}), 5)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
