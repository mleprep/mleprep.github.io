//go:build ignore

// 304. Range Sum Query 2D - Immutable (Medium)
// https://leetcode.com/problems/range-sum-query-2d-immutable/
// Run: go run 0304_range_sum_query_2d_immutable.go lc.go
package main

import (
	"fmt"
	"os"
)

type NumMatrix struct {
	P [][]int // P[r][c] = sum of the r x c top-left block
}

func Constructor(matrix [][]int) NumMatrix {
	R, C := len(matrix), len(matrix[0])
	P := make([][]int, R+1)
	for r := range P {
		P[r] = make([]int, C+1)
	}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			P[r+1][c+1] = matrix[r][c] + P[r][c+1] + P[r+1][c] - P[r][c]
		}
	}
	return NumMatrix{P: P}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	P := this.P // inclusion-exclusion on four prefix blocks
	return P[row2+1][col2+1] - P[row1][col2+1] - P[row2+1][col1] + P[row1][col1]
}

func main() {
	nm := Constructor([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})
	assertEq([]int{nm.SumRegion(2, 1, 4, 3), nm.SumRegion(1, 1, 2, 2), nm.SumRegion(1, 2, 2, 4)}, []int{8, 11, 12})
	small := Constructor([][]int{{1, 2}, {3, 4}})
	assertEq(small.SumRegion(0, 0, 1, 1), 10)
	assertEq(small.SumRegion(1, 1, 1, 1), 4)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
