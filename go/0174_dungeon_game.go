//go:build ignore

// 174. Dungeon Game (Hard)
// https://leetcode.com/problems/dungeon-game/
// Run: go run 0174_dungeon_game.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func calculateMinimumHP(dungeon [][]int) int {
	R, C := len(dungeon), len(dungeon[0])
	need := make([]int, C+1) // health needed on entering each cell, computed from the bottom right
	for c := range need {
		need[c] = math.MaxInt / 2
	}
	need[C-1] = 1 // virtual cell after the princess
	for r := R - 1; r >= 0; r-- {
		for c := C - 1; c >= 0; c-- {
			need[c] = max(1, min(need[c], need[c+1])-dungeon[r][c])
		}
	}
	return need[0]
}

func main() {
	assertEq(calculateMinimumHP([][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}), 7)
	assertEq(calculateMinimumHP([][]int{{0}}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
