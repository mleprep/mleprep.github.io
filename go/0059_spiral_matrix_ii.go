//go:build ignore

// 59. Spiral Matrix II (Medium)
// https://leetcode.com/problems/spiral-matrix-ii/
// Run: go run 0059_spiral_matrix_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func generateMatrix(n int) [][]int {
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}
	top, bottom, left, right, k := 0, n-1, 0, n-1, 1
	for top <= bottom {
		for c := left; c <= right; c++ {
			m[top][c] = k
			k++
		}
		for r := top + 1; r <= bottom; r++ {
			m[r][right] = k
			k++
		}
		if top < bottom {
			for c := right - 1; c >= left; c-- {
				m[bottom][c] = k
				k++
			}
		}
		if left < right {
			for r := bottom - 1; r > top; r-- {
				m[r][left] = k
				k++
			}
		}
		top, bottom, left, right = top+1, bottom-1, left+1, right-1
	}
	return m
}

func main() {
	assertEq(generateMatrix(3), [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}})
	assertEq(generateMatrix(1), [][]int{{1}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
