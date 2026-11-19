//go:build ignore

// 167. Two Sum II - Input Array Is Sorted (Medium)
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// Run: go run 0167_two_sum_ii_input_array_is_sorted.go lc.go
package main

import (
	"fmt"
	"os"
)

func twoSum(numbers []int, target int) []int {
	for l, r := 0, len(numbers)-1; l < r; {
		s := numbers[l] + numbers[r]
		if s == target {
			return []int{l + 1, r + 1}
		}
		if s < target {
			l++
		} else {
			r--
		}
	}
	return nil
}

func main() {
	assertEq(twoSum([]int{2, 7, 11, 15}, 9), []int{1, 2})
	assertEq(twoSum([]int{-1, 0}, -1), []int{1, 2})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
