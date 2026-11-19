//go:build ignore

// 1334. Find the City With the Smallest Number of Neighbors at a Threshold Distance (Medium)
// https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/
// Run: go run 1334_find_the_city_with_the_smallest_number_of_neighbors_at_a_threshold_distance.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	const inf = math.MaxInt32 / 2 // inf + inf must not overflow
	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, n)
		for j := range d[i] {
			if i != j {
				d[i][j] = inf
			}
		}
	}
	for _, e := range edges {
		d[e[0]][e[1]], d[e[1]][e[0]] = e[2], e[2]
	}
	for k := 0; k < n; k++ { // Floyd-Warshall
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if d[i][k]+d[k][j] < d[i][j] {
					d[i][j] = d[i][k] + d[k][j]
				}
			}
		}
	}
	best, city := math.MaxInt, -1
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if j != i && d[i][j] <= distanceThreshold {
				cnt++
			}
		}
		if cnt <= best { // ties go to the larger index
			best, city = cnt, i
		}
	}
	return city
}

func main() {
	assertEq(findTheCity(4, [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4), 3)
	assertEq(findTheCity(5, [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}, 2), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
