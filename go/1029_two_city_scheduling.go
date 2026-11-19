//go:build ignore

// 1029. Two City Scheduling (Medium)
// https://leetcode.com/problems/two-city-scheduling/
// Run: go run 1029_two_city_scheduling.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool { // most to gain from city A first
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})
	n, total := len(costs)/2, 0
	for i, c := range costs {
		if i < n {
			total += c[0]
		} else {
			total += c[1]
		}
	}
	return total
}

func main() {
	assertEq(twoCitySchedCost([][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}), 110)
	assertEq(twoCitySchedCost([][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}), 1859)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
