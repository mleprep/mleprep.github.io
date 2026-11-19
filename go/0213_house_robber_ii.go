//go:build ignore

// 213. House Robber II (Medium)
// https://leetcode.com/problems/house-robber-ii/
// Run: go run 0213_house_robber_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	line := func(arr []int) int {
		prev, cur := 0, 0
		for _, x := range arr {
			prev, cur = cur, max(cur, prev+x)
		}
		return cur
	}
	return max(line(nums[1:]), line(nums[:len(nums)-1])) // skip the first or the last house
}

func main() {
	assertEq(rob([]int{2, 3, 2}), 3)
	assertEq(rob([]int{1, 2, 3, 1}), 4)
	assertEq(rob([]int{1}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
