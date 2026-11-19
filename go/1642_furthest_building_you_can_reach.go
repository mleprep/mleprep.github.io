//go:build ignore

// 1642. Furthest Building You Can Reach (Medium)
// https://leetcode.com/problems/furthest-building-you-can-reach/
// Run: go run 1642_furthest_building_you_can_reach.go lc.go
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

func furthestBuilding(heights []int, bricks int, ladders int) int {
	h := &IntHeap{} // climbs covered by ladders: the largest ones
	for i := 0; i+1 < len(heights); i++ {
		d := heights[i+1] - heights[i]
		if d <= 0 {
			continue
		}
		heap.Push(h, d)
		if h.Len() > ladders { // the smallest climb pays with bricks
			bricks -= heap.Pop(h).(int)
			if bricks < 0 {
				return i
			}
		}
	}
	return len(heights) - 1
}

func main() {
	assertEq(furthestBuilding([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1), 4)
	assertEq(furthestBuilding([]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2), 7)
	assertEq(furthestBuilding([]int{14, 3, 19, 3}, 17, 0), 3)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
