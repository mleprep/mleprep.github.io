//go:build ignore

// 502. IPO (Hard)
// https://leetcode.com/problems/ipo/
// Run: go run 0502_ipo.go lc.go
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

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	idx := make([]int, len(profits))
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(a, b int) bool { return capital[idx[a]] < capital[idx[b]] })
	h, i := &MaxHeap{}, 0
	for ; k > 0; k-- {
		for i < len(idx) && capital[idx[i]] <= w { // newly affordable projects
			heap.Push(h, profits[idx[i]])
			i++
		}
		if h.Len() == 0 {
			break
		}
		w += heap.Pop(h).(int) // take the most profitable one
	}
	return w
}

func main() {
	assertEq(findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1}), 4)
	assertEq(findMaximizedCapital(3, 0, []int{1, 2, 3}, []int{0, 1, 2}), 6)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
