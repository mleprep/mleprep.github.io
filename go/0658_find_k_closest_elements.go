//go:build ignore

// 658. Find K Closest Elements (Medium)
// https://leetcode.com/problems/find-k-closest-elements/
// Run: go run 0658_find_k_closest_elements.go lc.go
package main

import (
	"fmt"
	"os"
)

func findClosestElements(arr []int, k int, x int) []int {
	lo, hi := 0, len(arr)-k // left end of the window
	for lo < hi {
		mid := (lo + hi) / 2
		if x-arr[mid] > arr[mid+k]-x { // arr[mid+k] is closer: move right
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return arr[lo : lo+k]
}

func main() {
	assertEq(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3), []int{1, 2, 3, 4})
	assertEq(findClosestElements([]int{1, 1, 2, 3, 4, 5}, 4, -1), []int{1, 1, 2, 3})
	assertEq(findClosestElements([]int{1, 3}, 1, 2), []int{1})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
