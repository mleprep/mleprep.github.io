//go:build ignore

// 96. Unique Binary Search Trees (Medium)
// https://leetcode.com/problems/unique-binary-search-trees/
// Run: go run 0096_unique_binary_search_trees.go lc.go
package main

import (
	"fmt"
	"os"
)

func numTrees(n int) int {
	G := make([]int, n+1) // G[m] = BSTs on m keys (Catalan numbers)
	G[0] = 1
	for m := 1; m <= n; m++ {
		for r := 1; r <= m; r++ { // root r: r-1 keys on the left, m-r on the right
			G[m] += G[r-1] * G[m-r]
		}
	}
	return G[n]
}

func main() {
	assertEq(numTrees(3), 5)
	assertEq(numTrees(1), 1)
	assertEq(numTrees(19), 1767263190)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
