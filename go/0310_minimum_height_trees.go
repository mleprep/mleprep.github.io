//go:build ignore

// 310. Minimum Height Trees (Medium)
// https://leetcode.com/problems/minimum-height-trees/
// Run: go run 0310_minimum_height_trees.go lc.go
package main

import (
	"fmt"
	"os"
)

func findMinHeightTrees(n int, edges [][]int) []int {
	if n <= 2 {
		out := []int{}
		for i := 0; i < n; i++ {
			out = append(out, i)
		}
		return out
	}
	adj := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
		deg[e[0]]++
		deg[e[1]]++
	}
	leaves := []int{}
	for i := 0; i < n; i++ {
		if deg[i] == 1 {
			leaves = append(leaves, i)
		}
	}
	remaining := n
	for remaining > 2 { // peel leaves layer by layer; the last 1-2 nodes are the centres
		remaining -= len(leaves)
		next := []int{}
		for _, leaf := range leaves {
			for _, nb := range adj[leaf] {
				deg[nb]--
				if deg[nb] == 1 {
					next = append(next, nb)
				}
			}
		}
		leaves = next
	}
	return leaves
}

func main() {
	assertUnordered(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}), []int{1})
	assertUnordered(findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}), []int{3, 4})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
