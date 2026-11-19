//go:build ignore

// 815. Bus Routes (Hard)
// https://leetcode.com/problems/bus-routes/
// Run: go run 0815_bus_routes.go lc.go
package main

import (
	"fmt"
	"os"
)

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	at := map[int][]int{} // stop -> routes through it
	for i, route := range routes {
		for _, s := range route {
			at[s] = append(at[s], i)
		}
	}
	used, seen, q := map[int]bool{}, map[int]bool{source: true}, []int{source}
	for buses := 1; len(q) > 0; buses++ { // BFS over stops, one level per bus taken
		next := []int{}
		for _, stop := range q {
			for _, r := range at[stop] {
				if used[r] {
					continue
				}
				used[r] = true
				for _, s := range routes[r] {
					if s == target {
						return buses
					}
					if !seen[s] {
						seen[s] = true
						next = append(next, s)
					}
				}
			}
		}
		q = next
	}
	return -1
}

func main() {
	assertEq(numBusesToDestination([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6), 2)
	assertEq(numBusesToDestination([][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12), -1)
	assertEq(numBusesToDestination([][]int{{1, 2}}, 1, 1), 0)
	assertEq(numBusesToDestination([][]int{{1, 2}}, 3, 1), -1)
	assertEq(numBusesToDestination([][]int{{1, 7}, {3, 5}}, 5, 5), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
