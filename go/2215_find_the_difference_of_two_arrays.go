//go:build ignore

// 2215. Find the Difference of Two Arrays (Easy)
// https://leetcode.com/problems/find-the-difference-of-two-arrays/
// Run: go run 2215_find_the_difference_of_two_arrays.go lc.go
package main

import (
	"fmt"
	"os"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	a, b := map[int]bool{}, map[int]bool{}
	for _, x := range nums1 {
		a[x] = true
	}
	for _, x := range nums2 {
		b[x] = true
	}
	minus := func(x, y map[int]bool) []int {
		out := []int{}
		for v := range x {
			if !y[v] {
				out = append(out, v)
			}
		}
		return out
	}
	return [][]int{minus(a, b), minus(b, a)}
}

func main() {
	r := findDifference([]int{1, 2, 3}, []int{2, 4, 6})
	assertUnordered(r[0], []int{1, 3})
	assertUnordered(r[1], []int{4, 6})
	r = findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2})
	assertEq(r, [][]int{{3}, {}})
	assertEq(findDifference([]int{-1}, []int{-1}), [][]int{{}, {}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
