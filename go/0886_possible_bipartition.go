//go:build ignore

// 886. Possible Bipartition (Medium)
// https://leetcode.com/problems/possible-bipartition/
// Run: go run 0886_possible_bipartition.go lc.go
package main

import (
	"fmt"
	"os"
)

func possibleBipartition(n int, dislikes [][]int) bool {
	g := make([][]int, n+1)
	for _, d := range dislikes {
		g[d[0]] = append(g[d[0]], d[1])
		g[d[1]] = append(g[d[1]], d[0])
	}
	color := make([]int, n+1) // 0 = uncolored, 1 / 2 = the two groups
	for s := 1; s <= n; s++ {
		if color[s] != 0 {
			continue
		}
		color[s] = 1
		q := []int{s}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for _, v := range g[u] {
				if color[v] == 0 {
					color[v] = 3 - color[u]
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
	assertEq(possibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 4}}), true)
	assertEq(possibleBipartition(3, [][]int{{1, 2}, {1, 3}, {2, 3}}), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
