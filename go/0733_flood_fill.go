//go:build ignore

// 733. Flood Fill (Easy)
// https://leetcode.com/problems/flood-fill/
// Run: go run 0733_flood_fill.go lc.go
package main

import (
	"fmt"
	"os"
)

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	old := image[sr][sc]
	if old == color {
		return image
	}
	R, C := len(image), len(image[0])
	image[sr][sc] = color
	st := [][2]int{{sr, sc}}
	for len(st) > 0 {
		p := st[len(st)-1]
		st = st[:len(st)-1]
		for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			nr, nc := p[0]+d[0], p[1]+d[1]
			if nr >= 0 && nr < R && nc >= 0 && nc < C && image[nr][nc] == old {
				image[nr][nc] = color
				st = append(st, [2]int{nr, nc})
			}
		}
	}
	return image
}

func main() {
	assertEq(floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2), [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}})
	assertEq(floodFill([][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0), [][]int{{0, 0, 0}, {0, 0, 0}})
	assertEq(floodFill([][]int{{0, 0, 0}, {0, 1, 1}}, 1, 1, 1), [][]int{{0, 0, 0}, {0, 1, 1}})
	assertEq(floodFill([][]int{{1}}, 0, 0, 3), [][]int{{3}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
