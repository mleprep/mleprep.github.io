//go:build ignore

// 1514. Path with Maximum Probability (Medium)
// https://leetcode.com/problems/path-with-maximum-probability/
// Run: go run 1514_path_with_maximum_probability.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"os"
)

type probItem struct {
	p float64
	u int
}

type probHeap []probItem // largest probability first

func (h probHeap) Len() int           { return len(h) }
func (h probHeap) Less(i, j int) bool { return h[i].p > h[j].p }
func (h probHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *probHeap) Push(x any)        { *h = append(*h, x.(probItem)) }
func (h *probHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	type edge struct {
		to int
		p  float64
	}
	g := make([][]edge, n)
	for i, e := range edges {
		g[e[0]] = append(g[e[0]], edge{e[1], succProb[i]})
		g[e[1]] = append(g[e[1]], edge{e[0], succProb[i]})
	}
	best := make([]float64, n)
	best[start_node] = 1
	pq := &probHeap{{1, start_node}}
	for pq.Len() > 0 { // Dijkstra on max-product: multiplying by p <= 1 never helps later
		cur := heap.Pop(pq).(probItem)
		if cur.u == end_node {
			return cur.p
		}
		if cur.p < best[cur.u] {
			continue
		}
		for _, e := range g[cur.u] {
			if np := cur.p * e.p; np > best[e.to] {
				best[e.to] = np
				heap.Push(pq, probItem{np, e.to})
			}
		}
	}
	return 0
}

func main() {
	assertEq(maxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.2}, 0, 2), 0.25)
	assertEq(maxProbability(3, [][]int{{0, 1}}, []float64{0.5}, 0, 2), 0.0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
