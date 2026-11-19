//go:build ignore

// 1732. Find the Highest Altitude (Easy)
// https://leetcode.com/problems/find-the-highest-altitude/
// Run: go run 1732_find_the_highest_altitude.go lc.go
package main

import (
	"fmt"
	"os"
)

func largestAltitude(gain []int) int {
	alt, best := 0, 0 // the start (altitude 0) counts too
	for _, g := range gain {
		alt += g
		best = max(best, alt)
	}
	return best
}

func main() {
	assertEq(largestAltitude([]int{-5, 1, 5, 0, -7}), 1)
	assertEq(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}), 0)
	assertEq(largestAltitude([]int{3}), 3)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
