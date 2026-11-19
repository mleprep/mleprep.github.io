//go:build ignore

// 695. Max Area of Island (Medium)
// https://leetcode.com/problems/max-area-of-island/
// Run: go run 0695_max_area_of_island.go lc.go
package main

import (
	"fmt"
	"os"
)

func maxAreaOfIsland(grid [][]int) int {
	R, C := len(grid), len(grid[0])
	var area func(r, c int) int
	area = func(r, c int) int {
		if r < 0 || r >= R || c < 0 || c >= C || grid[r][c] != 1 {
			return 0
		}
		grid[r][c] = 0 // visited
		return 1 + area(r+1, c) + area(r-1, c) + area(r, c+1) + area(r, c-1)
	}
	best := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			best = max(best, area(r, c))
		}
	}
	return best
}

func main() {
	assertEq(maxAreaOfIsland([][]int{{0, 0, 1, 0, 0}, {1, 1, 1, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 0}}), 4)
	assertEq(maxAreaOfIsland([][]int{{0, 0, 0, 0, 0}}), 0)
	assertEq(maxAreaOfIsland([][]int{{1}}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
