//go:build ignore

// 261. Graph Valid Tree (Medium)
// https://leetcode.com/problems/graph-valid-tree/
// Run: go run 0261_graph_valid_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 { // a tree has exactly n-1 edges
		return false
	}
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
	for _, e := range edges {
		ra, rb := find(e[0]), find(e[1])
		if ra == rb { // ...and no cycle
			return false
		}
		parent[ra] = rb
	}
	return true
}

func main() {
	assertEq(validTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}), true)
	assertEq(validTree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
