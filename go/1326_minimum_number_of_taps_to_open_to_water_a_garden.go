//go:build ignore

// 1326. Minimum Number of Taps to Open to Water a Garden (Hard)
// https://leetcode.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/
// Run: go run 1326_minimum_number_of_taps_to_open_to_water_a_garden.go lc.go
package main

import (
	"fmt"
	"os"
)

func minTaps(n int, ranges []int) int {
	reach := make([]int, n+1) // reach[lo] = farthest point a tap starting at lo covers
	for i, r := range ranges {
		lo, hi := max(0, i-r), min(n, i+r)
		reach[lo] = max(reach[lo], hi)
	}
	taps, end, far := 0, 0, 0 // then it is Jump Game II
	for i := 0; i < n; i++ {
		far = max(far, reach[i])
		if i == end {
			if far <= i {
				return -1 // a gap nothing covers
			}
			taps++
			end = far
		}
	}
	return taps
}

func main() {
	assertEq(minTaps(5, []int{3, 4, 1, 1, 0, 0}), 1)
	assertEq(minTaps(3, []int{0, 0, 0, 0}), -1)
	assertEq(minTaps(7, []int{1, 2, 1, 0, 2, 1, 0, 1}), 3)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
