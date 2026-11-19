//go:build ignore

// 1834. Single-Threaded CPU (Medium)
// https://leetcode.com/problems/single-threaded-cpu/
// Run: go run 1834_single_threaded_cpu.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type job struct{ proc, idx int }
type jobHeap []job // shortest processing time, then smallest index

func (h jobHeap) Len() int { return len(h) }
func (h jobHeap) Less(i, j int) bool {
	if h[i].proc != h[j].proc {
		return h[i].proc < h[j].proc
	}
	return h[i].idx < h[j].idx
}
func (h jobHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *jobHeap) Push(x any)   { *h = append(*h, x.(job)) }
func (h *jobHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func getOrder(tasks [][]int) []int {
	order := make([]int, len(tasks))
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(a, b int) bool { return tasks[order[a]][0] < tasks[order[b]][0] })
	h, res, time, i := &jobHeap{}, []int{}, 0, 0
	for len(res) < len(tasks) {
		if h.Len() == 0 && time < tasks[order[i]][0] { // idle until the next task arrives
			time = tasks[order[i]][0]
		}
		for i < len(order) && tasks[order[i]][0] <= time {
			heap.Push(h, job{tasks[order[i]][1], order[i]})
			i++
		}
		j := heap.Pop(h).(job)
		time += j.proc
		res = append(res, j.idx)
	}
	return res
}

func main() {
	assertEq(getOrder([][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}), []int{0, 2, 3, 1})
	assertEq(getOrder([][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}}), []int{4, 3, 2, 0, 1})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
