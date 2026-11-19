//go:build ignore

// 778. Swim in Rising Water (Hard)
// https://leetcode.com/problems/swim-in-rising-water/
// Run: go run 0778_swim_in_rising_water.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"os"
)

type cellHeap [][3]int // (time, row, col), smallest time first

func (h cellHeap) Len() int           { return len(h) }
func (h cellHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h cellHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *cellHeap) Push(x any)        { *h = append(*h, x.([3]int)) }
func (h *cellHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func swimInWater(grid [][]int) int {
	n := len(grid)
	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	seen[0][0] = true
	pq := &cellHeap{{grid[0][0], 0, 0}}
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for pq.Len() > 0 {
		cur := heap.Pop(pq).([3]int)
		t, r, c := cur[0], cur[1], cur[2]
		if r == n-1 && c == n-1 {
			return t
		}
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < n && nc >= 0 && nc < n && !seen[nr][nc] {
				seen[nr][nc] = true
				heap.Push(pq, [3]int{max(t, grid[nr][nc]), nr, nc})
			}
		}
	}
	return -1
}

func main() {
	assertEq(swimInWater([][]int{{0, 2}, {1, 3}}), 3)
	G := [][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}
	assertEq(swimInWater(G), 16)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
