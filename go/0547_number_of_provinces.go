//go:build ignore

// 547. Number of Provinces (Medium)
// https://leetcode.com/problems/number-of-provinces/
// Run: go run 0547_number_of_provinces.go lc.go
package main

import (
	"fmt"
	"os"
)

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	seen := make([]bool, n)
	var dfs func(u int)
	dfs = func(u int) {
		for v := 0; v < n; v++ {
			if isConnected[u][v] == 1 && !seen[v] {
				seen[v] = true
				dfs(v)
			}
		}
	}
	count := 0
	for u := 0; u < n; u++ {
		if !seen[u] {
			seen[u] = true
			dfs(u)
			count++
		}
	}
	return count
}

func main() {
	assertEq(findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}), 2)
	assertEq(findCircleNum([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}), 3)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
