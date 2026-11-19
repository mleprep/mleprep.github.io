//go:build ignore

// 1863. Sum of All Subset XOR Totals (Easy)
// https://leetcode.com/problems/sum-of-all-subset-xor-totals/
// Run: go run 1863_sum_of_all_subset_xor_totals.go lc.go
package main

import (
	"fmt"
	"os"
)

func subsetXORSum(nums []int) int {
	or := 0
	for _, x := range nums {
		or |= x
	}
	// each bit set in some number is set in exactly half of the 2^n subset XORs
	return or << (len(nums) - 1)
}

func main() {
	assertEq(subsetXORSum([]int{1, 3}), 6)
	assertEq(subsetXORSum([]int{5, 1, 6}), 28)
	assertEq(subsetXORSum([]int{3, 4, 5, 6, 7, 8}), 480)
	nums := []int{9, 2, 7, 12, 1}
	brute := 0
	for mask := 0; mask < 1<<len(nums); mask++ {
		x := 0
		for i, v := range nums {
			if mask>>i&1 == 1 {
				x ^= v
			}
		}
		brute += x
	}
	assertEq(subsetXORSum(nums), brute)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
