//go:build ignore

// 88. Merge Sorted Array (Easy)
// https://leetcode.com/problems/merge-sorted-array/
// Run: go run 0088_merge_sorted_array.go lc.go
package main

import (
	"fmt"
	"os"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1 // fill from the back
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	merge(a, 3, []int{2, 5, 6}, 3)
	assertEq(a, []int{1, 2, 2, 3, 5, 6})
	b := []int{0}
	merge(b, 0, []int{1}, 1)
	assertEq(b, []int{1})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
