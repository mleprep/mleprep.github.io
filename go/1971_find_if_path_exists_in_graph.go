//go:build ignore

// 1971. Find if Path Exists in Graph (Easy)
// https://leetcode.com/problems/find-if-path-exists-in-graph/
// Run: go run 1971_find_if_path_exists_in_graph.go lc.go
package main

import (
	"fmt"
	"os"
)

func validPath(n int, edges [][]int, source int, destination int) bool {
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
		parent[find(e[0])] = find(e[1])
	}
	return find(source) == find(destination)
}

func main() {
	assertEq(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2), true)
	assertEq(validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
