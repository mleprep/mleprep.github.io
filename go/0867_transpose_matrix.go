//go:build ignore

// 867. Transpose Matrix (Easy)
// https://leetcode.com/problems/transpose-matrix/
// Run: go run 0867_transpose_matrix.go lc.go
package main

import (
	"fmt"
	"os"
)

func transpose(matrix [][]int) [][]int {
	R, C := len(matrix), len(matrix[0])
	t := make([][]int, C)
	for c := range t {
		t[c] = make([]int, R)
		for r := 0; r < R; r++ {
			t[c][r] = matrix[r][c]
		}
	}
	return t
}

func main() {
	assertEq(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}), [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}})
	assertEq(transpose([][]int{{1, 2, 3}, {4, 5, 6}}), [][]int{{1, 4}, {2, 5}, {3, 6}})
	assertEq(transpose([][]int{{7}}), [][]int{{7}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
