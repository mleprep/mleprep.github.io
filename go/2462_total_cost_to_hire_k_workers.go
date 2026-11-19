//go:build ignore

// 2462. Total Cost to Hire K Workers (Medium)
// https://leetcode.com/problems/total-cost-to-hire-k-workers/
// Run: go run 2462_total_cost_to_hire_k_workers.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"os"
)

type IntHeap []int // min-heap

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func totalCost(costs []int, k int, candidates int) int64 {
	i, j := 0, len(costs)-1
	front, back := &IntHeap{}, &IntHeap{}
	total := 0
	for ; k > 0; k-- {
		for front.Len() < candidates && i <= j {
			heap.Push(front, costs[i])
			i++
		}
		for back.Len() < candidates && i <= j {
			heap.Push(back, costs[j])
			j--
		}
		if back.Len() == 0 || (front.Len() > 0 && (*front)[0] <= (*back)[0]) { // ties go to the front
			total += heap.Pop(front).(int)
		} else {
			total += heap.Pop(back).(int)
		}
	}
	return int64(total)
}

func main() {
	assertEq(totalCost([]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4), 11)
	assertEq(totalCost([]int{1, 2, 4, 1}, 3, 3), 4)
	assertEq(totalCost([]int{5}, 1, 1), 5)
	assertEq(totalCost([]int{2, 2, 2, 2}, 2, 1), 4)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
