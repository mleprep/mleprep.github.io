//go:build ignore

// 84. Largest Rectangle in Histogram (Hard)
// https://leetcode.com/problems/largest-rectangle-in-histogram/
// Run: go run 0084_largest_rectangle_in_histogram.go lc.go
package main

import (
	"fmt"
	"os"
)

func largestRectangleArea(heights []int) int {
	st, best := []int{}, 0 // increasing stack of indices
	for i := 0; i <= len(heights); i++ {
		h := 0 // sentinel bar after the end
		if i < len(heights) {
			h = heights[i]
		}
		for len(st) > 0 && heights[st[len(st)-1]] >= h {
			top := heights[st[len(st)-1]]
			st = st[:len(st)-1]
			left := -1
			if len(st) > 0 {
				left = st[len(st)-1]
			}
			best = max(best, top*(i-left-1))
		}
		st = append(st, i)
	}
	return best
}

func main() {
	assertEq(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}), 10)
	assertEq(largestRectangleArea([]int{2, 4}), 4)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
