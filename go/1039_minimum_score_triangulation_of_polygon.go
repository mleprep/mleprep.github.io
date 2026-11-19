//go:build ignore

// 1039. Minimum Score Triangulation of Polygon (Medium)
// https://leetcode.com/problems/minimum-score-triangulation-of-polygon/
// Run: go run 1039_minimum_score_triangulation_of_polygon.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n) // dp[i][j] = best triangulation of the polygon values[i..j]
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for gap := 2; gap < n; gap++ {
		for i := 0; i+gap < n; i++ {
			j := i + gap
			dp[i][j] = math.MaxInt
			for k := i + 1; k < j; k++ { // triangle (i, k, j) splits the rest in two
				dp[i][j] = min(dp[i][j], dp[i][k]+values[i]*values[k]*values[j]+dp[k][j])
			}
		}
	}
	return dp[0][n-1]
}

func main() {
	assertEq(minScoreTriangulation([]int{1, 2, 3}), 6)
	assertEq(minScoreTriangulation([]int{3, 7, 4, 5}), 144)
	assertEq(minScoreTriangulation([]int{1, 3, 1, 4, 1, 5}), 13)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
