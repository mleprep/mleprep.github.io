//go:build ignore

// 496. Next Greater Element I (Easy)
// https://leetcode.com/problems/next-greater-element-i/
// Run: go run 0496_next_greater_element_i.go lc.go
package main

import (
	"fmt"
	"os"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nxt, st := map[int]int{}, []int{}
	for _, x := range nums2 {
		for len(st) > 0 && st[len(st)-1] < x {
			nxt[st[len(st)-1]] = x
			st = st[:len(st)-1]
		}
		st = append(st, x)
	}
	res := make([]int, len(nums1))
	for i, x := range nums1 {
		if v, ok := nxt[x]; ok {
			res[i] = v
		} else {
			res[i] = -1
		}
	}
	return res
}

func main() {
	assertEq(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}), []int{-1, 3, -1})
	assertEq(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}), []int{3, -1})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
