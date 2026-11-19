//go:build ignore

// 463. Island Perimeter (Easy)
// https://leetcode.com/problems/island-perimeter/
// Run: go run 0463_island_perimeter.go lc.go
package main

import (
	"fmt"
	"os"
)

func islandPerimeter(grid [][]int) int {
	land, shared := 0, 0
	for r, row := range grid {
		for c, v := range row {
			if v == 1 {
				land++
				if r+1 < len(grid) && grid[r+1][c] == 1 { // each shared edge hides two sides
					shared++
				}
				if c+1 < len(row) && row[c+1] == 1 {
					shared++
				}
			}
		}
	}
	return 4*land - 2*shared
}

func main() {
	assertEq(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}), 16)
	assertEq(islandPerimeter([][]int{{1}}), 4)
	assertEq(islandPerimeter([][]int{{1, 0}}), 4)
	assertEq(islandPerimeter([][]int{{1, 1}, {1, 1}}), 8)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
