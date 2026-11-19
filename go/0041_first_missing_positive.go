//go:build ignore

// 41. First Missing Positive (Hard)
// https://leetcode.com/problems/first-missing-positive/
// Run: go run 0041_first_missing_positive.go lc.go
package main

import (
	"fmt"
	"os"
)

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		for nums[i] >= 1 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			j := nums[i] - 1 // put value v at index v-1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	assertEq(firstMissingPositive([]int{1, 2, 0}), 3)
	assertEq(firstMissingPositive([]int{3, 4, -1, 1}), 2)
	assertEq(firstMissingPositive([]int{7, 8, 9, 11, 12}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
