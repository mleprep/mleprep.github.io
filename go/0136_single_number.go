//go:build ignore

// 136. Single Number (Easy)
// https://leetcode.com/problems/single-number/
// Run: go run 0136_single_number.go lc.go
package main

import (
	"fmt"
	"os"
)

func singleNumber(nums []int) int {
	x := 0
	for _, v := range nums { // pairs cancel: a ^ a = 0
		x ^= v
	}
	return x
}

func main() {
	assertEq(singleNumber([]int{4, 1, 2, 1, 2}), 4)
	assertEq(singleNumber([]int{2, 2, 1}), 1)
	assertEq(singleNumber([]int{1}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
