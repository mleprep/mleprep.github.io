//go:build ignore

// 1192. Critical Connections in a Network (Hard)
// https://leetcode.com/problems/critical-connections-in-a-network/
// Run: go run 1192_critical_connections_in_a_network.go lc.go
package main

import (
	"fmt"
	"os"
)

func criticalConnections(n int, connections [][]int) [][]int {
	g := make([][]int, n)
	for _, c := range connections {
		g[c[0]] = append(g[c[0]], c[1])
		g[c[1]] = append(g[c[1]], c[0])
	}
	tin, low := make([]int, n), make([]int, n)
	for i := range tin {
		tin[i] = -1
	}
	res := [][]int{}
	timer := 0
	var dfs func(u, parent int)
	dfs = func(u, parent int) {
		tin[u], low[u] = timer, timer
		timer++
		for _, v := range g[u] {
			if v == parent {
				continue
			}
			if tin[v] == -1 {
				dfs(v, u)
				low[u] = min(low[u], low[v])
				if low[v] > tin[u] { // v's subtree cannot reach u or above without this edge
					res = append(res, []int{u, v})
				}
			} else {
				low[u] = min(low[u], tin[v])
			}
		}
	}
	dfs(0, -1)
	return res
}

func main() {
	assertUnorderedDeep(criticalConnections(4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}), [][]int{{1, 3}})
	assertUnorderedDeep(criticalConnections(2, [][]int{{0, 1}}), [][]int{{0, 1}})
	assertUnorderedDeep(criticalConnections(6, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}, {3, 4}, {4, 5}, {5, 3}}), [][]int{{1, 3}})
	assertUnorderedDeep(criticalConnections(4, [][]int{{0, 1}, {1, 2}, {2, 3}}), [][]int{{0, 1}, {1, 2}, {2, 3}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
