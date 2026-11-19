//go:build ignore

// 931. Minimum Falling Path Sum (Medium)
// https://leetcode.com/problems/minimum-falling-path-sum/
// Run: go run 0931_minimum_falling_path_sum.go lc.go
package main

import (
	"fmt"
	"os"
	"slices"
)

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	prev := append([]int(nil), matrix[0]...)
	for _, row := range matrix[1:] {
		cur := make([]int, n)
		for c := 0; c < n; c++ {
			best := prev[c]
			if c > 0 {
				best = min(best, prev[c-1])
			}
			if c+1 < n {
				best = min(best, prev[c+1])
			}
			cur[c] = row[c] + best
		}
		prev = cur
	}
	return slices.Min(prev)
}

func main() {
	assertEq(minFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}), 13)
	assertEq(minFallingPathSum([][]int{{-19, 57}, {-40, -5}}), -59)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
