//go:build ignore

// 295. Find Median from Data Stream (Hard)
// https://leetcode.com/problems/find-median-from-data-stream/
// Run: go run 0295_find_median_from_data_stream.go lc.go
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

type MedianFinder struct {
	lo IntHeap // smaller half, stored negated (a max-heap)
	hi IntHeap // larger half
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.lo, -num)
	heap.Push(&this.hi, -heap.Pop(&this.lo).(int)) // the largest of the small half moves up
	if this.hi.Len() > this.lo.Len() {
		heap.Push(&this.lo, -heap.Pop(&this.hi).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lo.Len() > this.hi.Len() {
		return float64(-this.lo[0])
	}
	return float64(-this.lo[0]+this.hi[0]) / 2
}

func main() {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	assertEq(mf.FindMedian(), 1.5)
	mf.AddNum(3)
	assertEq(mf.FindMedian(), 2.0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
