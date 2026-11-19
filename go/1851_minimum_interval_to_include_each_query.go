//go:build ignore

// 1851. Minimum Interval to Include Each Query (Hard)
// https://leetcode.com/problems/minimum-interval-to-include-each-query/
// Run: go run 1851_minimum_interval_to_include_each_query.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type sizeHeap [][2]int // (size, right end), smallest size first

func (h sizeHeap) Len() int           { return len(h) }
func (h sizeHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h sizeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *sizeHeap) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *sizeHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	order := make([]int, len(queries)) // answer queries in increasing order
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(a, b int) bool { return queries[order[a]] < queries[order[b]] })
	res := make([]int, len(queries))
	h, i := &sizeHeap{}, 0
	for _, qi := range order {
		q := queries[qi]
		for ; i < len(intervals) && intervals[i][0] <= q; i++ { // intervals starting by q
			heap.Push(h, [2]int{intervals[i][1] - intervals[i][0] + 1, intervals[i][1]})
		}
		for h.Len() > 0 && (*h)[0][1] < q { // drop those that ended before q
			heap.Pop(h)
		}
		res[qi] = -1
		if h.Len() > 0 {
			res[qi] = (*h)[0][0]
		}
	}
	return res
}

func main() {
	assertEq(minInterval([][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, []int{2, 3, 4, 5}), []int{3, 3, 1, 4})
	assertEq(minInterval([][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}}, []int{2, 19, 5, 22}), []int{2, -1, 4, 6})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
