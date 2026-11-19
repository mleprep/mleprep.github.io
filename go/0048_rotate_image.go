//go:build ignore

// 48. Rotate Image (Medium)
// https://leetcode.com/problems/rotate-image/
// Run: go run 0048_rotate_image.go lc.go
package main

import (
	"fmt"
	"os"
)

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ { // transpose...
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for _, row := range matrix { // ...then mirror each row
		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			row[l], row[r] = row[r], row[l]
		}
	}
}

func main() {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(m)
	assertEq(m, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}})
	m = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(m)
	assertEq(m, [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}})
	m = [][]int{{1}}
	rotate(m)
	assertEq(m, [][]int{{1}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
