//go:build ignore

// 62. Unique Paths (Medium)
// https://leetcode.com/problems/unique-paths/
// Run: go run 0062_unique_paths.go lc.go
package main

import (
	"fmt"
	"os"
)

func uniquePaths(m int, n int) int {
	row := make([]int, n) // paths to each cell of the current row
	for c := range row {
		row[c] = 1
	}
	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			row[c] += row[c-1] // from above (old value) + from the left
		}
	}
	return row[n-1]
}

func main() {
	assertEq(uniquePaths(3, 7), 28)
	assertEq(uniquePaths(3, 2), 3)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
