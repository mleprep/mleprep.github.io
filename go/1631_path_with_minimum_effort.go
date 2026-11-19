//go:build ignore

// 1631. Path With Minimum Effort (Medium)
// https://leetcode.com/problems/path-with-minimum-effort/
// Run: go run 1631_path_with_minimum_effort.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
)

type effortHeap [][3]int // (effort, row, col), smallest effort first

func (h effortHeap) Len() int           { return len(h) }
func (h effortHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h effortHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *effortHeap) Push(x any)        { *h = append(*h, x.([3]int)) }
func (h *effortHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func minimumEffortPath(heights [][]int) int {
	R, C := len(heights), len(heights[0])
	best := make([][]int, R)
	for i := range best {
		best[i] = make([]int, C)
		for j := range best[i] {
			best[i][j] = math.MaxInt
		}
	}
	best[0][0] = 0
	pq := &effortHeap{{0, 0, 0}}
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for pq.Len() > 0 {
		cur := heap.Pop(pq).([3]int)
		e, r, c := cur[0], cur[1], cur[2]
		if r == R-1 && c == C-1 {
			return e
		}
		if e > best[r][c] {
			continue // stale entry
		}
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < R && nc >= 0 && nc < C {
				diff := heights[nr][nc] - heights[r][c]
				if diff < 0 {
					diff = -diff
				}
				if ne := max(e, diff); ne < best[nr][nc] {
					best[nr][nc] = ne
					heap.Push(pq, [3]int{ne, nr, nc})
				}
			}
		}
	}
	return 0
}

func main() {
	assertEq(minimumEffortPath([][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}), 2)
	assertEq(minimumEffortPath([][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
