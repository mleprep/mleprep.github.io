//go:build ignore

// 63. Unique Paths II (Medium)
// https://leetcode.com/problems/unique-paths-ii/
// Run: go run 0063_unique_paths_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid[0])
	row := make([]int, n)
	row[0] = 1
	for _, r := range obstacleGrid {
		for c := 0; c < n; c++ {
			if r[c] == 1 {
				row[c] = 0
			} else if c > 0 {
				row[c] += row[c-1]
			}
		}
	}
	return row[n-1]
}

func main() {
	assertEq(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}), 2)
	assertEq(uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
