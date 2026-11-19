//go:build ignore

// 120. Triangle (Medium)
// https://leetcode.com/problems/triangle/
// Run: go run 0120_triangle.go lc.go
package main

import (
	"fmt"
	"os"
)

func minimumTotal(triangle [][]int) int {
	best := append([]int(nil), triangle[len(triangle)-1]...) // bottom-up: best path from each cell
	for r := len(triangle) - 2; r >= 0; r-- {
		for c := 0; c <= r; c++ {
			best[c] = triangle[r][c] + min(best[c], best[c+1])
		}
	}
	return best[0]
}

func main() {
	assertEq(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}), 11)
	assertEq(minimumTotal([][]int{{-10}}), -10)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
