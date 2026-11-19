//go:build ignore

// 73. Set Matrix Zeroes (Medium)
// https://leetcode.com/problems/set-matrix-zeroes/
// Run: go run 0073_set_matrix_zeroes.go lc.go
package main

import (
	"fmt"
	"os"
)

func setZeroes(matrix [][]int) {
	R, C := len(matrix), len(matrix[0])
	row0, col0 := false, false // the first row/column are reused as markers, so remember them
	for c := 0; c < C; c++ {
		row0 = row0 || matrix[0][c] == 0
	}
	for r := 0; r < R; r++ {
		col0 = col0 || matrix[r][0] == 0
	}
	for r := 1; r < R; r++ {
		for c := 1; c < C; c++ {
			if matrix[r][c] == 0 {
				matrix[r][0], matrix[0][c] = 0, 0
			}
		}
	}
	for r := 1; r < R; r++ {
		for c := 1; c < C; c++ {
			if matrix[r][0] == 0 || matrix[0][c] == 0 {
				matrix[r][c] = 0
			}
		}
	}
	if row0 {
		for c := 0; c < C; c++ {
			matrix[0][c] = 0
		}
	}
	if col0 {
		for r := 0; r < R; r++ {
			matrix[r][0] = 0
		}
	}
}

func main() {
	m := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(m)
	assertEq(m, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}})
	m = [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(m)
	assertEq(m, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
