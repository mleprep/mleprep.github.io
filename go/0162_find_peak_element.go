//go:build ignore

// 162. Find Peak Element (Medium)
// https://leetcode.com/problems/find-peak-element/
// Run: go run 0162_find_peak_element.go lc.go
package main

import (
	"fmt"
	"os"
)

func findPeakElement(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < nums[mid+1] { // uphill: a peak lies to the right
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func main() {
	assertEq(findPeakElement([]int{1, 2, 3, 1}), 2)
	p := findPeakElement([]int{1, 2, 1, 3, 5, 6, 4})
	assertEq(p == 1 || p == 5, true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
