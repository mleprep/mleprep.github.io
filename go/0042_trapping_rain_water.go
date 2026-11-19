//go:build ignore

// 42. Trapping Rain Water (Hard)
// https://leetcode.com/problems/trapping-rain-water/
// Run: go run 0042_trapping_rain_water.go lc.go
package main

import (
	"fmt"
	"os"
)

func trap(height []int) int {
	l, r := 0, len(height)-1
	lmax, rmax, water := 0, 0, 0
	for l < r {
		if height[l] < height[r] { // the left side is bounded by lmax
			lmax = max(lmax, height[l])
			water += lmax - height[l]
			l++
		} else {
			rmax = max(rmax, height[r])
			water += rmax - height[r]
			r--
		}
	}
	return water
}

func main() {
	assertEq(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), 6)
	assertEq(trap([]int{4, 2, 0, 3, 2, 5}), 9)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
