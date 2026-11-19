//go:build ignore

// 1466. Reorder Routes to Make All Paths Lead to the City Zero (Medium)
// https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/
// Run: go run 1466_reorder_routes_to_make_all_paths_lead_to_the_city_zero.go lc.go
package main

import (
	"fmt"
	"os"
)

func minReorder(n int, connections [][]int) int {
	g := make([][][2]int, n) // (neighbor, 1 if the road points away from city 0's side)
	for _, c := range connections {
		g[c[0]] = append(g[c[0]], [2]int{c[1], 1})
		g[c[1]] = append(g[c[1]], [2]int{c[0], 0})
	}
	seen, st, flips := make([]bool, n), []int{0}, 0
	seen[0] = true
	for len(st) > 0 {
		u := st[len(st)-1]
		st = st[:len(st)-1]
		for _, e := range g[u] {
			if !seen[e[0]] {
				seen[e[0]] = true
				flips += e[1]
				st = append(st, e[0])
			}
		}
	}
	return flips
}

func main() {
	assertEq(minReorder(6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}), 3)
	assertEq(minReorder(5, [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}), 2)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
