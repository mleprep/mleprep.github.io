//go:build ignore

// 740. Delete and Earn (Medium)
// https://leetcode.com/problems/delete-and-earn/
// Run: go run 0740_delete_and_earn.go lc.go
package main

import (
	"fmt"
	"os"
	"slices"
)

func deleteAndEarn(nums []int) int {
	gain := make([]int, slices.Max(nums)+1) // taking value v earns v * count(v)
	for _, x := range nums {
		gain[x] += x
	}
	prev, cur := 0, 0 // then it is House Robber over the values
	for _, g := range gain {
		prev, cur = cur, max(cur, prev+g)
	}
	return cur
}

func main() {
	assertEq(deleteAndEarn([]int{3, 4, 2}), 6)
	assertEq(deleteAndEarn([]int{2, 2, 3, 3, 3, 4}), 9)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
