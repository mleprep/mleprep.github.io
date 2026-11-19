//go:build ignore

// 1288. Remove Covered Intervals (Medium)
// https://leetcode.com/problems/remove-covered-intervals/
// Run: go run 1288_remove_covered_intervals.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { // start up; for equal starts the longest first
		a, b := intervals[i], intervals[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	keep, maxEnd := 0, -1
	for _, iv := range intervals { // covered exactly when an earlier interval reaches as far
		if iv[1] > maxEnd {
			keep++
			maxEnd = iv[1]
		}
	}
	return keep
}

func main() {
	assertEq(removeCoveredIntervals([][]int{{1, 4}, {3, 6}, {2, 8}}), 2)
	assertEq(removeCoveredIntervals([][]int{{1, 4}, {2, 3}}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
