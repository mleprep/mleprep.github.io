//go:build ignore

// 452. Minimum Number of Arrows to Burst Balloons (Medium)
// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/
// Run: go run 0452_minimum_number_of_arrows_to_burst_balloons.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	arrows, pos := 0, math.MinInt
	for _, p := range points { // shoot at the end of the first balloon not yet burst
		if p[0] > pos {
			arrows++
			pos = p[1]
		}
	}
	return arrows
}

func main() {
	assertEq(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}), 2)
	assertEq(findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}), 4)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
