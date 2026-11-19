//go:build ignore

// 215. Kth Largest Element in an Array (Medium)
// https://leetcode.com/problems/kth-largest-element-in-an-array/
// Run: go run 0215_kth_largest_element_in_an_array.go lc.go
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

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{} // the k largest seen so far
	for _, x := range nums {
		heap.Push(h, x)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

func main() {
	assertEq(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2), 5)
	assertEq(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4), 4)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
