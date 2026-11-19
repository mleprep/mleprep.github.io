//go:build ignore

// 80. Remove Duplicates from Sorted Array II (Medium)
// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
// Run: go run 0080_remove_duplicates_from_sorted_array_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func removeDuplicates(nums []int) int {
	k := 0
	for _, x := range nums {
		if k < 2 || x != nums[k-2] { // at most two copies
			nums[k] = x
			k++
		}
	}
	return k
}

func main() {
	a := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	k := removeDuplicates(a)
	assertEq(a[:k], []int{0, 0, 1, 1, 2, 3, 3})
	b := []int{1, 1, 1, 2, 2, 3}
	k = removeDuplicates(b)
	assertEq(b[:k], []int{1, 1, 2, 2, 3})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
