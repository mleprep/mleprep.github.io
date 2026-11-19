//go:build ignore

// 852. Peak Index in a Mountain Array (Medium)
// https://leetcode.com/problems/peak-index-in-a-mountain-array/
// Run: go run 0852_peak_index_in_a_mountain_array.go lc.go
package main

import (
	"fmt"
	"os"
)

func peakIndexInMountainArray(arr []int) int {
	lo, hi := 0, len(arr)-1
	for lo < hi { // first index where the array starts descending
		mid := lo + (hi-lo)/2
		if arr[mid] > arr[mid+1] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	assertEq(peakIndexInMountainArray([]int{0, 1, 0}), 1)
	assertEq(peakIndexInMountainArray([]int{0, 10, 5, 2}), 1)
	assertEq(peakIndexInMountainArray([]int{3, 5, 3, 2, 0}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
