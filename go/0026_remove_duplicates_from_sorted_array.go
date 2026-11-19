//go:build ignore

// 26. Remove Duplicates from Sorted Array (Easy)
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// Run: go run 0026_remove_duplicates_from_sorted_array.go lc.go
package main

import (
	"fmt"
	"os"
)

func removeDuplicates(nums []int) int {
	k := 0
	for _, x := range nums {
		if k == 0 || x != nums[k-1] {
			nums[k] = x
			k++
		}
	}
	return k
}

func main() {
	a := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(a)
	assertEq(a[:k], []int{0, 1, 2, 3, 4})
	b := []int{1, 1, 2}
	k = removeDuplicates(b)
	assertEq(b[:k], []int{1, 2})
	c := []int{7}
	k = removeDuplicates(c)
	assertEq(c[:k], []int{7})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
