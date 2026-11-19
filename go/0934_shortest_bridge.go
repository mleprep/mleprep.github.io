//go:build ignore

// 934. Shortest Bridge (Medium)
// https://leetcode.com/problems/shortest-bridge/
// Run: go run 0934_shortest_bridge.go lc.go
package main

import (
	"fmt"
	"os"
)

func shortestBridge(grid [][]int) int {
	n := len(grid)
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var start [2]int
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 1 {
				start = [2]int{r, c}
			}
		}
	}
	q, st := [][2]int{}, [][2]int{start}
	grid[start[0]][start[1]] = 2
	for len(st) > 0 { // mark one island as 2; it seeds the BFS
		p := st[len(st)-1]
		st = st[:len(st)-1]
		q = append(q, p)
		for _, d := range dirs {
			nr, nc := p[0]+d[0], p[1]+d[1]
			if nr >= 0 && nr < n && nc >= 0 && nc < n && grid[nr][nc] == 1 {
				grid[nr][nc] = 2
				st = append(st, [2]int{nr, nc})
			}
		}
	}
	for steps := 0; len(q) > 0; steps++ { // expand over water until the other island
		next := [][2]int{}
		for _, p := range q {
			for _, d := range dirs {
				nr, nc := p[0]+d[0], p[1]+d[1]
				if nr < 0 || nr >= n || nc < 0 || nc >= n {
					continue
				}
				if grid[nr][nc] == 1 {
					return steps
				}
				if grid[nr][nc] == 0 {
					grid[nr][nc] = 2
					next = append(next, [2]int{nr, nc})
				}
			}
		}
		q = next
	}
	return -1
}

func main() {
	assertEq(shortestBridge([][]int{{0, 1}, {1, 0}}), 1)
	assertEq(shortestBridge([][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}), 2)
	assertEq(shortestBridge([][]int{{1, 1, 1, 1, 1}, {1, 0, 0, 0, 1}, {1, 0, 1, 0, 1}, {1, 0, 0, 0, 1}, {1, 1, 1, 1, 1}}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
