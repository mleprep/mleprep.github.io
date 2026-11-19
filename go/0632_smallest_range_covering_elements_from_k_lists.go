//go:build ignore

// 632. Smallest Range Covering Elements from K Lists (Hard)
// https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/
// Run: go run 0632_smallest_range_covering_elements_from_k_lists.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
)

type entry struct{ v, i, j int } // value, list, position
type entryHeap []entry

func (h entryHeap) Len() int           { return len(h) }
func (h entryHeap) Less(i, j int) bool { return h[i].v < h[j].v }
func (h entryHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *entryHeap) Push(x any)        { *h = append(*h, x.(entry)) }
func (h *entryHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func smallestRange(nums [][]int) []int {
	h, hi := &entryHeap{}, math.MinInt
	for i, row := range nums { // one pointer per list; the range is [heap min, hi]
		heap.Push(h, entry{row[0], i, 0})
		hi = max(hi, row[0])
	}
	best := []int{(*h)[0].v, hi}
	for {
		e := heap.Pop(h).(entry)
		if hi-e.v < best[1]-best[0] {
			best = []int{e.v, hi}
		}
		if e.j+1 == len(nums[e.i]) { // a list is exhausted
			return best
		}
		next := nums[e.i][e.j+1]
		hi = max(hi, next)
		heap.Push(h, entry{next, e.i, e.j + 1})
	}
}

func main() {
	assertEq(smallestRange([][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}), []int{20, 24})
	assertEq(smallestRange([][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}), []int{1, 1})
	assertEq(smallestRange([][]int{{1}, {2}, {3}}), []int{1, 3})
	assertEq(smallestRange([][]int{{10, 10}, {11, 11}}), []int{10, 11})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
