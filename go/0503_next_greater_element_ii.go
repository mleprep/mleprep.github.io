//go:build ignore

// 503. Next Greater Element II (Medium)
// https://leetcode.com/problems/next-greater-element-ii/
// Run: go run 0503_next_greater_element_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res, st := make([]int, n), []int{}
	for i := range res {
		res[i] = -1
	}
	for i := 0; i < 2*n; i++ { // two passes handle the wrap-around
		x := nums[i%n]
		for len(st) > 0 && nums[st[len(st)-1]] < x {
			res[st[len(st)-1]] = x
			st = st[:len(st)-1]
		}
		if i < n {
			st = append(st, i)
		}
	}
	return res
}

func main() {
	assertEq(nextGreaterElements([]int{1, 2, 1}), []int{2, -1, 2})
	assertEq(nextGreaterElements([]int{1, 2, 3, 4, 3}), []int{2, 3, 4, -1, 4})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
