//go:build ignore

// 1584. Min Cost to Connect All Points (Medium)
// https://leetcode.com/problems/min-cost-to-connect-all-points/
// Run: go run 1584_min_cost_to_connect_all_points.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	dist := make([]int, n) // cheapest edge from the tree to each point
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[0] = 0
	used := make([]bool, n)
	total := 0
	for step := 0; step < n; step++ { // dense Prim: O(n^2)
		u := -1
		for i := 0; i < n; i++ {
			if !used[i] && (u == -1 || dist[i] < dist[u]) {
				u = i
			}
		}
		used[u] = true
		total += dist[u]
		for v := 0; v < n; v++ {
			if !used[v] {
				dist[v] = min(dist[v], abs(points[v][0]-points[u][0])+abs(points[v][1]-points[u][1]))
			}
		}
	}
	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	assertEq(minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}), 20)
	assertEq(minCostConnectPoints([][]int{{3, 12}, {-2, 5}, {-4, 1}}), 18)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
