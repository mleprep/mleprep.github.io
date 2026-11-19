//go:build ignore

// 802. Find Eventual Safe States (Medium)
// https://leetcode.com/problems/find-eventual-safe-states/
// Run: go run 0802_find_eventual_safe_states.go lc.go
package main

import (
	"fmt"
	"os"
)

func eventualSafeNodes(graph [][]int) []int {
	color := make([]int, len(graph)) // 0 unvisited, 1 on the path or unsafe, 2 safe
	var safe func(u int) bool
	safe = func(u int) bool {
		if color[u] != 0 {
			return color[u] == 2
		}
		color[u] = 1
		for _, v := range graph[u] {
			if !safe(v) {
				return false // stays 1: unsafe
			}
		}
		color[u] = 2
		return true
	}
	res := []int{}
	for u := range graph {
		if safe(u) {
			res = append(res, u)
		}
	}
	return res
}

func main() {
	assertEq(eventualSafeNodes([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}), []int{2, 4, 5, 6})
	assertEq(eventualSafeNodes([][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}), []int{4})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
