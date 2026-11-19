//go:build ignore

// 630. Course Schedule III (Hard)
// https://leetcode.com/problems/course-schedule-iii/
// Run: go run 0630_course_schedule_iii.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool { return courses[i][1] < courses[j][1] }) // by deadline
	h, t := &MaxHeap{}, 0
	for _, c := range courses {
		heap.Push(h, c[0])
		t += c[0]
		if t > c[1] { // over the deadline: drop the longest course taken
			t -= heap.Pop(h).(int)
		}
	}
	return h.Len()
}

func main() {
	assertEq(scheduleCourse([][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}), 3)
	assertEq(scheduleCourse([][]int{{1, 2}}), 1)
	assertEq(scheduleCourse([][]int{{3, 2}, {4, 3}}), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
