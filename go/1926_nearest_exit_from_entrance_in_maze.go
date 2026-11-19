//go:build ignore

// 1926. Nearest Exit from Entrance in Maze (Medium)
// https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/
// Run: go run 1926_nearest_exit_from_entrance_in_maze.go lc.go
package main

import (
	"fmt"
	"os"
)

func nearestExit(maze [][]byte, entrance []int) int {
	R, C := len(maze), len(maze[0])
	start := [2]int{entrance[0], entrance[1]}
	seen, q := map[[2]int]bool{start: true}, [][2]int{start}
	for steps := 1; len(q) > 0; steps++ {
		next := [][2]int{}
		for _, p := range q {
			for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				nr, nc := p[0]+d[0], p[1]+d[1]
				if nr < 0 || nr >= R || nc < 0 || nc >= C || maze[nr][nc] != '.' || seen[[2]int{nr, nc}] {
					continue
				}
				if nr == 0 || nr == R-1 || nc == 0 || nc == C-1 { // a border cell that is not the entrance
					return steps
				}
				seen[[2]int{nr, nc}] = true
				next = append(next, [2]int{nr, nc})
			}
		}
		q = next
	}
	return -1
}

func main() {
	assertEq(nearestExit(byteGrid("++.+", "...+", "+++."), []int{1, 2}), 1)
	assertEq(nearestExit(byteGrid("+++", "...", "+++"), []int{1, 0}), 2)
	assertEq(nearestExit(byteGrid(".+"), []int{0, 0}), -1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
