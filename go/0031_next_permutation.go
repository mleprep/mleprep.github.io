//go:build ignore

// 31. Next Permutation (Medium)
// https://leetcode.com/problems/next-permutation/
// Run: go run 0031_next_permutation.go lc.go
package main

import (
	"fmt"
	"os"
)

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	for l, r := i+1, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}

func main() {
	for _, tc := range [][2][]int{{{1, 2, 3}, {1, 3, 2}}, {{3, 2, 1}, {1, 2, 3}}, {{1, 1, 5}, {1, 5, 1}}, {{1, 3, 5, 4, 2}, {1, 4, 2, 3, 5}}} {
		nextPermutation(tc[0])
		assertEq(tc[0], tc[1])
	}
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
