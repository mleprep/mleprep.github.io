//go:build ignore

// 417. Pacific Atlantic Water Flow (Medium)
// https://leetcode.com/problems/pacific-atlantic-water-flow/
// Run: go run 0417_pacific_atlantic_water_flow.go lc.go
package main

import (
	"fmt"
	"os"
)

func pacificAtlantic(heights [][]int) [][]int {
	R, C := len(heights), len(heights[0])
	reach := func(starts [][2]int) [][]bool { // cells that can flow into this ocean (walk uphill from it)
		seen := make([][]bool, R)
		for i := range seen {
			seen[i] = make([]bool, C)
		}
		st := starts
		for _, p := range starts {
			seen[p[0]][p[1]] = true
		}
		for len(st) > 0 {
			p := st[len(st)-1]
			st = st[:len(st)-1]
			for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				nr, nc := p[0]+d[0], p[1]+d[1]
				if nr >= 0 && nr < R && nc >= 0 && nc < C && !seen[nr][nc] && heights[nr][nc] >= heights[p[0]][p[1]] {
					seen[nr][nc] = true
					st = append(st, [2]int{nr, nc})
				}
			}
		}
		return seen
	}
	var pacStart, atlStart [][2]int
	for c := 0; c < C; c++ {
		pacStart = append(pacStart, [2]int{0, c})
		atlStart = append(atlStart, [2]int{R - 1, c})
	}
	for r := 0; r < R; r++ {
		pacStart = append(pacStart, [2]int{r, 0})
		atlStart = append(atlStart, [2]int{r, C - 1})
	}
	pac, atl := reach(pacStart), reach(atlStart)
	res := [][]int{}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if pac[r][c] && atl[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}
	return res
}

func main() {
	H := [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}
	assertUnordered(pacificAtlantic(H), [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}})
	assertEq(pacificAtlantic([][]int{{1}}), [][]int{{0, 0}})
	assertUnordered(pacificAtlantic([][]int{{2, 1}, {1, 2}}), [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
