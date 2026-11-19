//go:build ignore

// 685. Redundant Connection II (Hard)
// https://leetcode.com/problems/redundant-connection-ii/
// Run: go run 0685_redundant_connection_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges)
	parentOf := make([]int, n+1) // 0 = no parent yet (nodes are 1..n)
	var cand1, cand2 []int       // the two edges into a node with two parents
	for _, e := range edges {
		u, v := e[0], e[1]
		if parentOf[v] != 0 {
			cand1, cand2 = []int{parentOf[v], v}, []int{u, v}
		} else {
			parentOf[v] = u
		}
	}
	root := make([]int, n+1)
	for i := range root {
		root[i] = i
	}
	find := func(x int) int {
		for root[x] != x {
			root[x] = root[root[x]]
			x = root[x]
		}
		return x
	}
	for _, e := range edges {
		u, v := e[0], e[1]
		if cand2 != nil && u == cand2[0] && v == cand2[1] {
			continue // try the tree without the later candidate
		}
		ru, rv := find(u), find(v)
		if ru == rv { // a cycle even so
			if cand1 != nil {
				return cand1
			}
			return []int{u, v}
		}
		root[rv] = ru
	}
	return cand2
}

func main() {
	assertEq(findRedundantDirectedConnection([][]int{{1, 2}, {1, 3}, {2, 3}}), []int{2, 3})
	assertEq(findRedundantDirectedConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}), []int{4, 1})
	assertEq(findRedundantDirectedConnection([][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}}), []int{2, 1})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
