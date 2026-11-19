//go:build ignore

// 57. Insert Interval (Medium)
// https://leetcode.com/problems/insert-interval/
// Run: go run 0057_insert_interval.go lc.go
package main

import (
	"fmt"
	"os"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	res, i, n := [][]int{}, 0, len(intervals)
	s, e := newInterval[0], newInterval[1]
	for ; i < n && intervals[i][1] < s; i++ { // entirely before
		res = append(res, intervals[i])
	}
	for ; i < n && intervals[i][0] <= e; i++ { // overlapping: absorb
		s, e = min(s, intervals[i][0]), max(e, intervals[i][1])
	}
	res = append(res, []int{s, e})
	return append(res, intervals[i:]...) // entirely after
}

func main() {
	assertEq(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}), [][]int{{1, 5}, {6, 9}})
	assertEq(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}), [][]int{{1, 2}, {3, 10}, {12, 16}})
	assertEq(insert([][]int{}, []int{5, 7}), [][]int{{5, 7}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
