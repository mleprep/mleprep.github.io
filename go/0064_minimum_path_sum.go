//go:build ignore

// 64. Minimum Path Sum (Medium)
// https://leetcode.com/problems/minimum-path-sum/
// Run: go run 0064_minimum_path_sum.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func minPathSum(grid [][]int) int {
	n := len(grid[0])
	row := make([]int, n) // cheapest path sum to each cell of the current row
	for c := 1; c < n; c++ {
		row[c] = math.MaxInt / 2 // nothing above the first row
	}
	for _, r := range grid {
		row[0] += r[0]
		for c := 1; c < n; c++ {
			row[c] = r[c] + min(row[c], row[c-1])
		}
	}
	return row[n-1]
}

func main() {
	assertEq(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}), 7)
	assertEq(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}), 12)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
