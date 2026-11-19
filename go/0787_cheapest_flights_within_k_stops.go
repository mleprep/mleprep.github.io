//go:build ignore

// 787. Cheapest Flights Within K Stops (Medium)
// https://leetcode.com/problems/cheapest-flights-within-k-stops/
// Run: go run 0787_cheapest_flights_within_k_stops.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const inf = math.MaxInt32
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[src] = 0
	for round := 0; round <= k; round++ { // Bellman-Ford limited to k+1 edges
		next := append([]int(nil), dist...) // relax from last round's values only
		for _, f := range flights {
			if dist[f[0]] != inf && dist[f[0]]+f[2] < next[f[1]] {
				next[f[1]] = dist[f[0]] + f[2]
			}
		}
		dist = next
	}
	if dist[dst] == inf {
		return -1
	}
	return dist[dst]
}

func main() {
	F := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
	assertEq(findCheapestPrice(4, F, 0, 3, 1), 700)
	assertEq(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0), 500)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
