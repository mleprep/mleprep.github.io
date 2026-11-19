//go:build ignore

// 456. 132 Pattern (Medium)
// https://leetcode.com/problems/132-pattern/
// Run: go run 0456_132_pattern.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func find132pattern(nums []int) bool {
	st, two := []int{}, math.MinInt
	for i := len(nums) - 1; i >= 0; i-- {
		x := nums[i]
		if x < two {
			return true
		}
		for len(st) > 0 && st[len(st)-1] < x {
			two = st[len(st)-1] // x acts as "3", the popped value as "2"
			st = st[:len(st)-1]
		}
		st = append(st, x)
	}
	return false
}

func main() {
	assertEq(find132pattern([]int{1, 2, 3, 4}), false)
	assertEq(find132pattern([]int{3, 1, 4, 2}), true)
	assertEq(find132pattern([]int{-1, 3, 2, 0}), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
