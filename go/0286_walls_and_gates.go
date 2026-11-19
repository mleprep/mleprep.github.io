//go:build ignore

// 286. Walls and Gates (Medium)
// https://leetcode.com/problems/walls-and-gates/
// Run: go run 0286_walls_and_gates.go lc.go
package main

import (
	"fmt"
	"os"
)

func wallsAndGates(rooms [][]int) {
	const INF = 2147483647
	R, C := len(rooms), len(rooms[0])
	q := [][2]int{}
	for r := 0; r < R; r++ { // BFS from every gate at once
		for c := 0; c < C; c++ {
			if rooms[r][c] == 0 {
				q = append(q, [2]int{r, c})
			}
		}
	}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			nr, nc := p[0]+d[0], p[1]+d[1]
			if nr >= 0 && nr < R && nc >= 0 && nc < C && rooms[nr][nc] == INF {
				rooms[nr][nc] = rooms[p[0]][p[1]] + 1
				q = append(q, [2]int{nr, nc})
			}
		}
	}
}

func main() {
	const I = 2147483647
	g := [][]int{{I, -1, 0, I}, {I, I, I, -1}, {I, -1, I, -1}, {0, -1, I, I}}
	wallsAndGates(g)
	assertEq(g, [][]int{{3, -1, 0, 1}, {2, 2, 1, -1}, {1, -1, 2, -1}, {0, -1, 3, 4}})
	g2 := [][]int{{-1}}
	wallsAndGates(g2)
	assertEq(g2, [][]int{{-1}})
	g3 := [][]int{{I, 0}}
	wallsAndGates(g3)
	assertEq(g3, [][]int{{1, 0}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
