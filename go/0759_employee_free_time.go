//go:build ignore

// 759. Employee Free Time (Hard)
// https://leetcode.com/problems/employee-free-time/
// Run: go run 0759_employee_free_time.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

type Interval struct { // provided by LeetCode
	Start int
	End   int
}

func employeeFreeTime(schedule [][]*Interval) []*Interval {
	ivs := []*Interval{}
	for _, emp := range schedule { // whose interval it is does not matter
		ivs = append(ivs, emp...)
	}
	sort.Slice(ivs, func(i, j int) bool { return ivs[i].Start < ivs[j].Start })
	res := []*Interval{}
	end := ivs[0].End // furthest busy time so far
	for _, iv := range ivs[1:] {
		if iv.Start > end {
			res = append(res, &Interval{Start: end, End: iv.Start})
		}
		end = max(end, iv.End)
	}
	return res
}

func main() {
	free := func(s [][][]int) [][]int {
		sched := [][]*Interval{}
		for _, emp := range s {
			row := []*Interval{}
			for _, iv := range emp {
				row = append(row, &Interval{Start: iv[0], End: iv[1]})
			}
			sched = append(sched, row)
		}
		out := [][]int{}
		for _, iv := range employeeFreeTime(sched) {
			out = append(out, []int{iv.Start, iv.End})
		}
		return out
	}
	assertEq(free([][][]int{{{1, 2}, {5, 6}}, {{1, 3}}, {{4, 10}}}), [][]int{{3, 4}})
	assertEq(free([][][]int{{{1, 3}, {6, 7}}, {{2, 4}}, {{2, 5}, {9, 12}}}), [][]int{{5, 6}, {7, 9}})
	assertEq(free([][][]int{{{1, 2}}, {{2, 3}}}), [][]int{})
	assertEq(free([][][]int{{{1, 2}, {8, 9}}, {{3, 4}}}), [][]int{{2, 3}, {4, 8}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
