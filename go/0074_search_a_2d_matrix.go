//go:build ignore

// 74. Search a 2D Matrix (Medium)
// https://leetcode.com/problems/search-a-2d-matrix/
// Run: go run 0074_search_a_2d_matrix.go lc.go
package main

import (
	"fmt"
	"os"
)

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	lo, hi := 0, m*n-1 // binary search the flattened matrix
	for lo <= hi {
		mid := lo + (hi-lo)/2
		v := matrix[mid/n][mid%n]
		switch {
		case v == target:
			return true
		case v < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return false
}

func main() {
	M := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	assertEq(searchMatrix(M, 3), true)
	assertEq(searchMatrix(M, 13), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
