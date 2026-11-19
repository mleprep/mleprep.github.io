//go:build ignore

// 218. The Skyline Problem (Hard)
// https://leetcode.com/problems/the-skyline-problem/
// Run: go run 0218_the_skyline_problem.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
)

type heightHeap [][2]int // (height, right edge), tallest first

func (h heightHeap) Len() int           { return len(h) }
func (h heightHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h heightHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *heightHeap) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *heightHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func getSkyline(buildings [][]int) [][]int {
	xs := []int{} // every edge is a candidate key point
	for _, b := range buildings {
		xs = append(xs, b[0], b[1])
	}
	sort.Ints(xs)
	res := [][]int{}
	h := &heightHeap{{0, math.MaxInt}} // the ground never ends
	i := 0
	for k, x := range xs {
		if k > 0 && x == xs[k-1] {
			continue
		}
		for ; i < len(buildings) && buildings[i][0] <= x; i++ { // buildings started by x
			heap.Push(h, [2]int{buildings[i][2], buildings[i][1]})
		}
		for (*h)[0][1] <= x { // lazily drop the tallest while it has ended
			heap.Pop(h)
		}
		if top := (*h)[0][0]; len(res) == 0 || res[len(res)-1][1] != top {
			res = append(res, []int{x, top})
		}
	}
	return res
}

func main() {
	B := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	assertEq(getSkyline(B), [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}})
	assertEq(getSkyline([][]int{{0, 2, 3}, {2, 5, 3}}), [][]int{{0, 3}, {5, 0}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
