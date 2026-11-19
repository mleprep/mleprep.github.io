//go:build ignore

// 240. Search a 2D Matrix II (Medium)
// https://leetcode.com/problems/search-a-2d-matrix-ii/
// Run: go run 0240_search_a_2d_matrix_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func searchMatrix(matrix [][]int, target int) bool {
	r, c := 0, len(matrix[0])-1 // start top-right: left is smaller, down is bigger
	for r < len(matrix) && c >= 0 {
		v := matrix[r][c]
		switch {
		case v == target:
			return true
		case v > target:
			c--
		default:
			r++
		}
	}
	return false
}

func main() {
	M := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	assertEq(searchMatrix(M, 5), true)
	assertEq(searchMatrix(M, 20), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
