//go:build ignore

// 75. Sort Colors (Medium)
// https://leetcode.com/problems/sort-colors/
// Run: go run 0075_sort_colors.go lc.go
package main

import (
	"fmt"
	"os"
)

func sortColors(nums []int) {
	lo, i, hi := 0, 0, len(nums)-1 // [0,lo) = 0s, [lo,i) = 1s, (hi,end] = 2s
	for i <= hi {
		switch nums[i] {
		case 0:
			nums[lo], nums[i] = nums[i], nums[lo]
			lo++
			i++
		case 2:
			nums[hi], nums[i] = nums[i], nums[hi]
			hi--
		default:
			i++
		}
	}
}

func main() {
	a := []int{2, 0, 2, 1, 1, 0}
	sortColors(a)
	assertEq(a, []int{0, 0, 1, 1, 2, 2})
	b := []int{2, 0, 1}
	sortColors(b)
	assertEq(b, []int{0, 1, 2})
	c := []int{1, 1, 0, 0, 2, 2, 1}
	sortColors(c)
	assertEq(c, []int{0, 0, 1, 1, 1, 2, 2})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
