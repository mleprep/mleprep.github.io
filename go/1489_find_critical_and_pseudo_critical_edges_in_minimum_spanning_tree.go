//go:build ignore

// 1489. Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree (Hard)
// https://leetcode.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/
// Run: go run 1489_find_critical_and_pseudo_critical_edges_in_minimum_spanning_tree.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	order := make([]int, len(edges))
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(a, b int) bool { return edges[order[a]][2] < edges[order[b]][2] })
	mst := func(skip, force int) int { // Kruskal without edge `skip` / starting with edge `force`
		parent := make([]int, n)
		for i := range parent {
			parent[i] = i
		}
		find := func(x int) int {
			for parent[x] != x {
				parent[x] = parent[parent[x]]
				x = parent[x]
			}
			return x
		}
		weight, used := 0, 0
		if force >= 0 {
			e := edges[force]
			parent[find(e[0])] = find(e[1])
			weight, used = e[2], 1
		}
		for _, i := range order {
			e := edges[i]
			if i != skip && find(e[0]) != find(e[1]) {
				parent[find(e[0])] = find(e[1])
				weight += e[2]
				used++
			}
		}
		if used < n-1 {
			return math.MaxInt // disconnected
		}
		return weight
	}
	base := mst(-1, -1)
	critical, pseudo := []int{}, []int{}
	for i := range edges {
		if mst(i, -1) > base {
			critical = append(critical, i)
		} else if mst(-1, i) == base {
			pseudo = append(pseudo, i)
		}
	}
	return [][]int{critical, pseudo}
}

func main() {
	assertEq(findCriticalAndPseudoCriticalEdges(5, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6}}), [][]int{{0, 1}, {2, 3, 4, 5}})
	assertEq(findCriticalAndPseudoCriticalEdges(4, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 1}}), [][]int{{}, {0, 1, 2, 3}})
	assertEq(findCriticalAndPseudoCriticalEdges(3, [][]int{{0, 1, 1}, {1, 2, 2}, {0, 2, 3}}), [][]int{{0, 1}, {}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
