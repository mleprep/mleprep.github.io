//go:build ignore

// 1679. Max Number of K-Sum Pairs (Medium)
// https://leetcode.com/problems/max-number-of-k-sum-pairs/
// Run: go run 1679_max_number_of_k_sum_pairs.go lc.go
package main

import (
	"fmt"
	"os"
)

func maxOperations(nums []int, k int) int {
	waiting, ops := map[int]int{}, 0
	for _, x := range nums {
		if waiting[k-x] > 0 { // pair x with an earlier k-x
			waiting[k-x]--
			ops++
		} else {
			waiting[x]++
		}
	}
	return ops
}

func main() {
	assertEq(maxOperations([]int{1, 2, 3, 4}, 5), 2)
	assertEq(maxOperations([]int{3, 1, 3, 4, 3}, 6), 1)
	assertEq(maxOperations([]int{2, 2, 2, 2}, 4), 2)
	assertEq(maxOperations([]int{1}, 2), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
