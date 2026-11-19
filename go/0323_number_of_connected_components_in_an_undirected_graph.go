//go:build ignore

// 323. Number of Connected Components in an Undirected Graph (Medium)
// https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/
// Run: go run 0323_number_of_connected_components_in_an_undirected_graph.go lc.go
package main

import (
	"fmt"
	"os"
)

func countComponents(n int, edges [][]int) int {
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
	comps := n
	for _, e := range edges {
		if ra, rb := find(e[0]), find(e[1]); ra != rb {
			parent[ra] = rb
			comps-- // each successful union merges two components
		}
	}
	return comps
}

func main() {
	assertEq(countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}}), 2)
	assertEq(countComponents(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
