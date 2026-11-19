//go:build ignore

// 994. Rotting Oranges (Medium)
// https://leetcode.com/problems/rotting-oranges/
// Run: go run 0994_rotting_oranges.go lc.go
package main

import (
	"fmt"
	"os"
)

func orangesRotting(grid [][]int) int {
	R, C := len(grid), len(grid[0])
	q, fresh := [][2]int{}, 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if grid[r][c] == 2 {
				q = append(q, [2]int{r, c})
			} else if grid[r][c] == 1 {
				fresh++
			}
		}
	}
	minutes := 0
	for len(q) > 0 && fresh > 0 { // multi-source BFS, one level per minute
		minutes++
		next := [][2]int{}
		for _, p := range q {
			for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				nr, nc := p[0]+d[0], p[1]+d[1]
				if nr >= 0 && nr < R && nc >= 0 && nc < C && grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					fresh--
					next = append(next, [2]int{nr, nc})
				}
			}
		}
		q = next
	}
	if fresh > 0 {
		return -1
	}
	return minutes
}

func main() {
	assertEq(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}), 4)
	assertEq(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}), -1)
	assertEq(orangesRotting([][]int{{0, 2}}), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
