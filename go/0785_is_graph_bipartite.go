//go:build ignore

// 785. Is Graph Bipartite? (Medium)
// https://leetcode.com/problems/is-graph-bipartite/
// Run: go run 0785_is_graph_bipartite.go lc.go
package main

import (
	"fmt"
	"os"
)

func isBipartite(graph [][]int) bool {
	color := make([]int, len(graph))
	for i := range color {
		color[i] = -1
	}
	for s := range graph { // the graph may be disconnected
		if color[s] != -1 {
			continue
		}
		color[s] = 0
		q := []int{s}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for _, v := range graph[u] {
				if color[v] == -1 {
					color[v] = 1 - color[u]
					q = append(q, v)
				} else if color[v] == color[u] {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	assertEq(isBipartite([][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}), false)
	assertEq(isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
