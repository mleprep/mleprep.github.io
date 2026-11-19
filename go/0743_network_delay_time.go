//go:build ignore

// 743. Network Delay Time (Medium)
// https://leetcode.com/problems/network-delay-time/
// Run: go run 0743_network_delay_time.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"os"
)

type distHeap [][2]int // (distance, node), smallest distance first

func (h distHeap) Len() int           { return len(h) }
func (h distHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h distHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *distHeap) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *distHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
	g := make([][][2]int, n+1) // u -> (v, w)
	for _, t := range times {
		g[t[0]] = append(g[t[0]], [2]int{t[1], t[2]})
	}
	dist := map[int]int{} // finalized nodes
	pq := &distHeap{{0, k}}
	for pq.Len() > 0 {
		top := heap.Pop(pq).([2]int)
		d, u := top[0], top[1]
		if _, done := dist[u]; done {
			continue
		}
		dist[u] = d
		for _, e := range g[u] {
			if _, done := dist[e[0]]; !done {
				heap.Push(pq, [2]int{d + e[1], e[0]})
			}
		}
	}
	if len(dist) < n {
		return -1
	}
	best := 0
	for _, d := range dist {
		best = max(best, d)
	}
	return best
}

func main() {
	assertEq(networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2), 2)
	assertEq(networkDelayTime([][]int{{1, 2, 1}}, 2, 2), -1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
