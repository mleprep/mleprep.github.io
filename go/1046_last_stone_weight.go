//go:build ignore

// 1046. Last Stone Weight (Easy)
// https://leetcode.com/problems/last-stone-weight/
// Run: go run 1046_last_stone_weight.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"os"
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

func lastStoneWeight(stones []int) int {
	h := MaxHeap(append([]int{}, stones...))
	heap.Init(&h)
	for h.Len() > 1 { // smash the two heaviest
		a, b := heap.Pop(&h).(int), heap.Pop(&h).(int)
		if a != b {
			heap.Push(&h, a-b)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return h[0]
}

func main() {
	assertEq(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}), 1)
	assertEq(lastStoneWeight([]int{1}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
