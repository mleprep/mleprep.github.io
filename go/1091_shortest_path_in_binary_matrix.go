//go:build ignore

// 1091. Shortest Path in Binary Matrix (Medium)
// https://leetcode.com/problems/shortest-path-in-binary-matrix/
// Run: go run 1091_shortest_path_in_binary_matrix.go lc.go
package main

import (
	"fmt"
	"os"
)

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	q := [][3]int{{0, 0, 1}} // (r, c, path length)
	grid[0][0] = 1
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		if p[0] == n-1 && p[1] == n-1 {
			return p[2]
		}
		for dr := -1; dr <= 1; dr++ { // 8 directions
			for dc := -1; dc <= 1; dc++ {
				nr, nc := p[0]+dr, p[1]+dc
				if nr >= 0 && nr < n && nc >= 0 && nc < n && grid[nr][nc] == 0 {
					grid[nr][nc] = 1
					q = append(q, [3]int{nr, nc, p[2] + 1})
				}
			}
		}
	}
	return -1
}

func main() {
	assertEq(shortestPathBinaryMatrix([][]int{{0, 1}, {1, 0}}), 2)
	assertEq(shortestPathBinaryMatrix([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}), 4)
	assertEq(shortestPathBinaryMatrix([][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}), -1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
