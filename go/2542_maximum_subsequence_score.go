//go:build ignore

// 2542. Maximum Subsequence Score (Medium)
// https://leetcode.com/problems/maximum-subsequence-score/
// Run: go run 2542_maximum_subsequence_score.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"os"
	"sort"
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

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	idx := make([]int, len(nums1))
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(a, b int) bool { return nums2[idx[a]] > nums2[idx[b]] }) // by nums2 desc
	h, total, best := &IntHeap{}, 0, 0
	for _, i := range idx { // nums2[i] is the minimum of the chosen set so far
		heap.Push(h, nums1[i])
		total += nums1[i]
		if h.Len() > k { // keep the k largest nums1 values
			total -= heap.Pop(h).(int)
		}
		if h.Len() == k {
			best = max(best, total*nums2[i])
		}
	}
	return int64(best)
}

func main() {
	assertEq(maxScore([]int{1, 3, 3, 2}, []int{2, 1, 3, 4}, 3), 12)
	assertEq(maxScore([]int{4, 2, 3, 1, 1}, []int{7, 5, 10, 9, 6}, 1), 30)
	assertEq(maxScore([]int{2, 1, 14, 12}, []int{11, 7, 13, 6}, 3), 168)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
