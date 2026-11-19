//go:build ignore

// 1730. Shortest Path to Get Food (Medium)
// https://leetcode.com/problems/shortest-path-to-get-food/
// Run: go run 1730_shortest_path_to_get_food.go lc.go
package main

import (
	"fmt"
	"os"
)

func getFood(grid [][]byte) int {
	R, C := len(grid), len(grid[0])
	var start [2]int
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if grid[r][c] == '*' {
				start = [2]int{r, c}
			}
		}
	}
	seen, q := map[[2]int]bool{start: true}, [][2]int{start}
	for steps := 1; len(q) > 0; steps++ {
		next := [][2]int{}
		for _, p := range q {
			for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				nb := [2]int{p[0] + d[0], p[1] + d[1]}
				if nb[0] < 0 || nb[0] >= R || nb[1] < 0 || nb[1] >= C || grid[nb[0]][nb[1]] == 'X' || seen[nb] {
					continue
				}
				if grid[nb[0]][nb[1]] == '#' {
					return steps
				}
				seen[nb] = true
				next = append(next, nb)
			}
		}
		q = next
	}
	return -1
}

func main() {
	assertEq(getFood(byteGrid("XXXXXX", "X*OOOX", "XOO#OX", "XXXXXX")), 3)
	assertEq(getFood(byteGrid("XXXXX", "X*XOX", "XOX#X", "XXXXX")), -1)
	assertEq(getFood(byteGrid("XXXXXXXX", "X*OXO#OX", "XOOXOOXX", "XOOOO#OX", "XXXXXXXX")), 6)
	assertEq(getFood(byteGrid("*#")), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
