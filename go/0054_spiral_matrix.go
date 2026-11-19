//go:build ignore

// 54. Spiral Matrix (Medium)
// https://leetcode.com/problems/spiral-matrix/
// Run: go run 0054_spiral_matrix.go lc.go
package main

import (
	"fmt"
	"os"
)

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for top <= bottom && left <= right { // peel one ring per round
		for c := left; c <= right; c++ {
			res = append(res, matrix[top][c])
		}
		for r := top + 1; r <= bottom; r++ {
			res = append(res, matrix[r][right])
		}
		if top < bottom && left < right { // a single row or column has no way back
			for c := right - 1; c >= left; c-- {
				res = append(res, matrix[bottom][c])
			}
			for r := bottom - 1; r > top; r-- {
				res = append(res, matrix[r][left])
			}
		}
		top, bottom, left, right = top+1, bottom-1, left+1, right-1
	}
	return res
}

func main() {
	assertEq(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}), []int{1, 2, 3, 6, 9, 8, 7, 4, 5})
	assertEq(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}), []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
