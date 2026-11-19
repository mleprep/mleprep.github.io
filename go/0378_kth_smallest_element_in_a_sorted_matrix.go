//go:build ignore

// 378. Kth Smallest Element in a Sorted Matrix (Medium)
// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/
// Run: go run 0378_kth_smallest_element_in_a_sorted_matrix.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"os"
)

type cell struct{ v, r, c int }
type cellHeap []cell

func (h cellHeap) Len() int           { return len(h) }
func (h cellHeap) Less(i, j int) bool { return h[i].v < h[j].v }
func (h cellHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *cellHeap) Push(x any)        { *h = append(*h, x.(cell)) }
func (h *cellHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	h := &cellHeap{}
	for r := 0; r < min(n, k); r++ { // the first column of each row
		heap.Push(h, cell{matrix[r][0], r, 0})
	}
	for i := 0; i < k-1; i++ { // pop k-1 smallest, advancing along their rows
		c := heap.Pop(h).(cell)
		if c.c+1 < n {
			heap.Push(h, cell{matrix[c.r][c.c+1], c.r, c.c + 1})
		}
	}
	return (*h)[0].v
}

func main() {
	assertEq(kthSmallest([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8), 13)
	assertEq(kthSmallest([][]int{{-5}}, 1), -5)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
