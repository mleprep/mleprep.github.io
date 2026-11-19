//go:build ignore

// 11. Container With Most Water (Medium)
// https://leetcode.com/problems/container-with-most-water/
// Run: go run 0011_container_with_most_water.go lc.go
package main

import (
	"fmt"
	"os"
)

func maxArea(height []int) int {
	best := 0
	for l, r := 0, len(height)-1; l < r; {
		best = max(best, min(height[l], height[r])*(r-l))
		if height[l] < height[r] { // the shorter wall can only get better by moving
			l++
		} else {
			r--
		}
	}
	return best
}

func main() {
	assertEq(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)
	assertEq(maxArea([]int{1, 1}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
