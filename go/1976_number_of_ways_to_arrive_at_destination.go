//go:build ignore

// 1976. Number of Ways to Arrive at Destination (Medium)
// https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/
// Run: go run 1976_number_of_ways_to_arrive_at_destination.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
)

type pathHeap [][2]int // (distance, node), smallest distance first

func (h pathHeap) Len() int           { return len(h) }
func (h pathHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h pathHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *pathHeap) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *pathHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func countPaths(n int, roads [][]int) int {
	const mod = 1_000_000_007
	g := make([][][2]int, n)
	for _, r := range roads {
		g[r[0]] = append(g[r[0]], [2]int{r[1], r[2]})
		g[r[1]] = append(g[r[1]], [2]int{r[0], r[2]})
	}
	dist, ways := make([]int, n), make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[0], ways[0] = 0, 1
	pq := &pathHeap{{0, 0}}
	for pq.Len() > 0 {
		cur := heap.Pop(pq).([2]int)
		d, u := cur[0], cur[1]
		if d > dist[u] {
			continue
		}
		for _, e := range g[u] {
			v, t := e[0], e[1]
			if d+t < dist[v] { // strictly shorter: the shortest ways to v all go through u
				dist[v], ways[v] = d+t, ways[u]
				heap.Push(pq, [2]int{dist[v], v})
			} else if d+t == dist[v] {
				ways[v] = (ways[v] + ways[u]) % mod
			}
		}
	}
	return ways[n-1]
}

func main() {
	R := [][]int{{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3}, {6, 3, 3}, {3, 5, 1}, {6, 5, 1}, {2, 5, 1}, {0, 4, 5}, {4, 6, 2}}
	assertEq(countPaths(7, R), 4)
	assertEq(countPaths(2, [][]int{{1, 0, 10}}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
