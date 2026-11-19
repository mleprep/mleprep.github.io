//go:build ignore

// 853. Car Fleet (Medium)
// https://leetcode.com/problems/car-fleet/
// Run: go run 0853_car_fleet.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func carFleet(target int, position []int, speed []int) int {
	idx := make([]int, len(position))
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(a, b int) bool { return position[idx[a]] > position[idx[b]] }) // closest to target first
	fleets, slowest := 0, 0.0
	for _, i := range idx {
		t := float64(target-position[i]) / float64(speed[i])
		if t > slowest { // can't catch the fleet ahead
			fleets++
			slowest = t
		}
	}
	return fleets
}

func main() {
	assertEq(carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}), 3)
	assertEq(carFleet(100, []int{0, 2, 4}, []int{4, 2, 1}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
