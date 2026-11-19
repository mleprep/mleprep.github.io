//go:build ignore

// 329. Longest Increasing Path in a Matrix (Hard)
// https://leetcode.com/problems/longest-increasing-path-in-a-matrix/
// Run: go run 0329_longest_increasing_path_in_a_matrix.go lc.go
package main

import (
	"fmt"
	"os"
)

func longestIncreasingPath(matrix [][]int) int {
	R, C := len(matrix), len(matrix[0])
	memo := make([][]int, R) // longest path starting at (r, c); 0 = not computed
	for i := range memo {
		memo[i] = make([]int, C)
	}
	var best func(r, c int) int
	best = func(r, c int) int {
		if memo[r][c] > 0 {
			return memo[r][c]
		}
		res := 1
		for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < R && nc >= 0 && nc < C && matrix[nr][nc] > matrix[r][c] {
				res = max(res, 1+best(nr, nc))
			}
		}
		memo[r][c] = res
		return res
	}
	ans := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			ans = max(ans, best(r, c))
		}
	}
	return ans
}

func main() {
	assertEq(longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}), 4)
	assertEq(longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}), 4)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
