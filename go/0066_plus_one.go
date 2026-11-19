//go:build ignore

// 66. Plus One (Easy)
// https://leetcode.com/problems/plus-one/
// Run: go run 0066_plus_one.go lc.go
package main

import (
	"fmt"
	"os"
)

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0 // 9 + 1: carry on
	}
	return append([]int{1}, digits...)
}

func main() {
	assertEq(plusOne([]int{1, 2, 3}), []int{1, 2, 4})
	assertEq(plusOne([]int{9, 9}), []int{1, 0, 0})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
