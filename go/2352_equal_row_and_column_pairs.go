//go:build ignore

// 2352. Equal Row and Column Pairs (Medium)
// https://leetcode.com/problems/equal-row-and-column-pairs/
// Run: go run 2352_equal_row_and_column_pairs.go lc.go
package main

import (
	"fmt"
	"os"
)

func equalPairs(grid [][]int) int {
	n := len(grid)
	rows := map[string]int{}
	for _, r := range grid {
		rows[fmt.Sprint(r)]++
	}
	res := 0
	for c := 0; c < n; c++ {
		col := make([]int, n)
		for r := 0; r < n; r++ {
			col[r] = grid[r][c]
		}
		res += rows[fmt.Sprint(col)]
	}
	return res
}

func main() {
	assertEq(equalPairs([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}), 1)
	assertEq(equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}), 3)
	assertEq(equalPairs([][]int{{1}}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
