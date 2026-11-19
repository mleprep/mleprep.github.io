//go:build ignore

// 200. Number of Islands (Medium)
// https://leetcode.com/problems/number-of-islands/
// Run: go run 0200_number_of_islands.go lc.go
package main

import (
	"fmt"
	"os"
)

func numIslands(grid [][]byte) int {
	R, C := len(grid), len(grid[0])
	sink := func(r, c int) { // flood-fill one island with an explicit stack
		st := [][2]int{{r, c}}
		grid[r][c] = '0'
		for len(st) > 0 {
			p := st[len(st)-1]
			st = st[:len(st)-1]
			for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				nr, nc := p[0]+d[0], p[1]+d[1]
				if nr >= 0 && nr < R && nc >= 0 && nc < C && grid[nr][nc] == '1' {
					grid[nr][nc] = '0'
					st = append(st, [2]int{nr, nc})
				}
			}
		}
	}
	count := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if grid[r][c] == '1' {
				count++
				sink(r, c)
			}
		}
	}
	return count
}

func main() {
	assertEq(numIslands(byteGrid("11000", "11000", "00100", "00011")), 3)
	assertEq(numIslands(byteGrid("0")), 0)
	assertEq(numIslands(byteGrid("11110", "11010", "11000", "00000")), 1)
	assertEq(numIslands(byteGrid("101", "010", "101")), 5)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
