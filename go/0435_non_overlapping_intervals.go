//go:build ignore

// 435. Non-overlapping Intervals (Medium)
// https://leetcode.com/problems/non-overlapping-intervals/
// Run: go run 0435_non_overlapping_intervals.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	keep, end := 0, math.MinInt
	for _, iv := range intervals { // keep the interval that ends first, as often as possible
		if iv[0] >= end {
			keep++
			end = iv[1]
		}
	}
	return len(intervals) - keep
}

func main() {
	assertEq(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}), 1)
	assertEq(eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}), 2)
	assertEq(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}}), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
