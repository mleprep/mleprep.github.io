//go:build ignore

// 268. Missing Number (Easy)
// https://leetcode.com/problems/missing-number/
// Run: go run 0268_missing_number.go lc.go
package main

import (
	"fmt"
	"os"
)

func missingNumber(nums []int) int {
	x := len(nums)
	for i, v := range nums { // xor all indices 0..n and all values: only the missing one survives
		x ^= i ^ v
	}
	return x
}

func main() {
	assertEq(missingNumber([]int{3, 0, 1}), 2)
	assertEq(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}), 8)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
