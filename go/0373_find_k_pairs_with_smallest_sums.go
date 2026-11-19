//go:build ignore

// 373. Find K Pairs with Smallest Sums (Medium)
// https://leetcode.com/problems/find-k-pairs-with-smallest-sums/
// Run: go run 0373_find_k_pairs_with_smallest_sums.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"os"
)

type pair struct{ sum, i, j int }
type pairHeap []pair

func (h pairHeap) Len() int           { return len(h) }
func (h pairHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h pairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *pairHeap) Push(x any)        { *h = append(*h, x.(pair)) }
func (h *pairHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := &pairHeap{}
	for i := 0; i < min(k, len(nums1)); i++ { // (i, 0) for every row
		heap.Push(h, pair{nums1[i] + nums2[0], i, 0})
	}
	res := [][]int{}
	for h.Len() > 0 && len(res) < k {
		p := heap.Pop(h).(pair)
		res = append(res, []int{nums1[p.i], nums2[p.j]})
		if p.j+1 < len(nums2) {
			heap.Push(h, pair{nums1[p.i] + nums2[p.j+1], p.i, p.j + 1})
		}
	}
	return res
}

func main() {
	assertUnordered(kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3), [][]int{{1, 2}, {1, 4}, {1, 6}})
	assertUnordered(kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 2), [][]int{{1, 1}, {1, 1}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
