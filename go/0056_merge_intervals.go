//go:build ignore

// 56. Merge Intervals (Medium)
// https://leetcode.com/problems/merge-intervals/
// Run: go run 0056_merge_intervals.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	res := [][]int{}
	for _, iv := range intervals {
		if n := len(res); n > 0 && iv[0] <= res[n-1][1] { // overlaps the last merged interval
			res[n-1][1] = max(res[n-1][1], iv[1])
		} else {
			res = append(res, []int{iv[0], iv[1]})
		}
	}
	return res
}

func main() {
	assertEq(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}), [][]int{{1, 6}, {8, 10}, {15, 18}})
	assertEq(merge([][]int{{1, 4}, {4, 5}}), [][]int{{1, 5}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
