//go:build ignore

// 542. 01 Matrix (Medium)
// https://leetcode.com/problems/01-matrix/
// Run: go run 0542_01_matrix.go lc.go
package main

import (
	"fmt"
	"os"
)

func updateMatrix(mat [][]int) [][]int {
	R, C := len(mat), len(mat[0])
	dist := make([][]int, R)
	q := [][2]int{}
	for r := range dist {
		dist[r] = make([]int, C)
		for c := range dist[r] {
			if mat[r][c] == 0 {
				q = append(q, [2]int{r, c}) // BFS from every 0 at once
			} else {
				dist[r][c] = -1
			}
		}
	}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			nr, nc := p[0]+d[0], p[1]+d[1]
			if nr >= 0 && nr < R && nc >= 0 && nc < C && dist[nr][nc] == -1 {
				dist[nr][nc] = dist[p[0]][p[1]] + 1
				q = append(q, [2]int{nr, nc})
			}
		}
	}
	return dist
}

func main() {
	assertEq(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}), [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}})
	assertEq(updateMatrix([][]int{{0}, {1}, {1}}), [][]int{{0}, {1}, {2}})
	assertEq(updateMatrix([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 0}}), [][]int{{4, 3, 2}, {3, 2, 1}, {2, 1, 0}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
