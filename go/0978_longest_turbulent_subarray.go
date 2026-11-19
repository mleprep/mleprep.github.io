//go:build ignore

// 978. Longest Turbulent Subarray (Medium)
// https://leetcode.com/problems/longest-turbulent-subarray/
// Run: go run 0978_longest_turbulent_subarray.go lc.go
package main

import (
	"fmt"
	"os"
)

func maxTurbulenceSize(arr []int) int {
	best, up, down := 1, 1, 1 // longest run ending here going up / down
	for i := 1; i < len(arr); i++ {
		switch {
		case arr[i] > arr[i-1]:
			up, down = down+1, 1
		case arr[i] < arr[i-1]:
			up, down = 1, up+1
		default:
			up, down = 1, 1
		}
		best = max(best, up, down)
	}
	return best
}

func main() {
	assertEq(maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9}), 5)
	assertEq(maxTurbulenceSize([]int{4, 8, 12, 16}), 2)
	assertEq(maxTurbulenceSize([]int{100}), 1)
	assertEq(maxTurbulenceSize([]int{9, 9}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
