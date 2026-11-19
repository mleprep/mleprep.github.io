//go:build ignore

// 973. K Closest Points to Origin (Medium)
// https://leetcode.com/problems/k-closest-points-to-origin/
// Run: go run 0973_k_closest_points_to_origin.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"os"
)

type point struct{ d, x, y int }
type pointHeap []point // max-heap by distance

func (h pointHeap) Len() int           { return len(h) }
func (h pointHeap) Less(i, j int) bool { return h[i].d > h[j].d }
func (h pointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *pointHeap) Push(x any)        { *h = append(*h, x.(point)) }
func (h *pointHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := &pointHeap{}
	for _, p := range points {
		heap.Push(h, point{p[0]*p[0] + p[1]*p[1], p[0], p[1]})
		if h.Len() > k { // drop the farthest
			heap.Pop(h)
		}
	}
	res := [][]int{}
	for _, p := range *h {
		res = append(res, []int{p.x, p.y})
	}
	return res
}

func main() {
	assertEq(kClosest([][]int{{1, 3}, {-2, 2}}, 1), [][]int{{-2, 2}})
	assertUnordered(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2), [][]int{{-2, 4}, {3, 3}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
