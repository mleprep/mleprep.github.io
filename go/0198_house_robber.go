//go:build ignore

// 198. House Robber (Medium)
// https://leetcode.com/problems/house-robber/
// Run: go run 0198_house_robber.go lc.go
package main

import (
	"fmt"
	"os"
)

func rob(nums []int) int {
	prev, cur := 0, 0 // best up to i-2, best up to i-1
	for _, x := range nums {
		prev, cur = cur, max(cur, prev+x)
	}
	return cur
}

func main() {
	assertEq(rob([]int{1, 2, 3, 1}), 4)
	assertEq(rob([]int{2, 7, 9, 3, 1}), 12)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
