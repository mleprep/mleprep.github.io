//go:build ignore

// 684. Redundant Connection (Medium)
// https://leetcode.com/problems/redundant-connection/
// Run: go run 0684_redundant_connection.go lc.go
package main

import (
	"fmt"
	"os"
)

func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
	}
	find := func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]] // path halving
			x = parent[x]
		}
		return x
	}
	for _, e := range edges {
		ra, rb := find(e[0]), find(e[1])
		if ra == rb { // already connected: this edge closes the cycle
			return e
		}
		parent[ra] = rb
	}
	return nil
}

func main() {
	assertEq(findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}), []int{2, 3})
	assertEq(findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}), []int{1, 4})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
