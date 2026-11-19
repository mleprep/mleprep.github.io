//go:build ignore

// 260. Single Number III (Medium)
// https://leetcode.com/problems/single-number-iii/
// Run: go run 0260_single_number_iii.go lc.go
package main

import (
	"fmt"
	"os"
)

func singleNumber(nums []int) []int {
	x := 0
	for _, v := range nums {
		x ^= v // = a ^ b
	}
	diff := x & -x // a bit where a and b differ splits the numbers into two groups
	a := 0
	for _, v := range nums {
		if v&diff != 0 {
			a ^= v
		}
	}
	return []int{a, x ^ a}
}

func main() {
	assertUnordered(singleNumber([]int{1, 2, 1, 3, 2, 5}), []int{3, 5})
	assertUnordered(singleNumber([]int{-1, 0}), []int{-1, 0})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
