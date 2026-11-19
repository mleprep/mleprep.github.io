//go:build ignore

// 703. Kth Largest Element in a Stream (Easy)
// https://leetcode.com/problems/kth-largest-element-in-a-stream/
// Run: go run 0703_kth_largest_element_in_a_stream.go lc.go
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

type KthLargest struct {
	k int
	h IntHeap // the k largest values; the root is the k-th largest
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, x := range nums {
		kl.Add(x)
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.h, val)
	if this.h.Len() > this.k {
		heap.Pop(&this.h)
	}
	return this.h[0]
}

func main() {
	kl := Constructor(3, []int{4, 5, 8, 2})
	got := []int{}
	for _, v := range []int{3, 5, 10, 9, 4} {
		got = append(got, kl.Add(v))
	}
	assertEq(got, []int{4, 5, 5, 8, 8})
	kl2 := Constructor(1, []int{})
	got2 := []int{}
	for _, v := range []int{-3, -2, -4, 0, 4} {
		got2 = append(got2, kl2.Add(v))
	}
	assertEq(got2, []int{-3, -2, -2, 0, 4})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
