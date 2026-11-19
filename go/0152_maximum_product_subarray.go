//go:build ignore

// 152. Maximum Product Subarray (Medium)
// https://leetcode.com/problems/maximum-product-subarray/
// Run: go run 0152_maximum_product_subarray.go lc.go
package main

import (
	"fmt"
	"os"
)

func maxProduct(nums []int) int {
	best, hi, lo := nums[0], nums[0], nums[0] // largest and smallest product ending here
	for _, x := range nums[1:] {
		a, b := hi*x, lo*x // a negative x swaps the roles of hi and lo
		hi, lo = max(x, a, b), min(x, a, b)
		best = max(best, hi)
	}
	return best
}

func main() {
	assertEq(maxProduct([]int{2, 3, -2, 4}), 6)
	assertEq(maxProduct([]int{-2, 0, -1}), 0)
	assertEq(maxProduct([]int{-2, 3, -4}), 24)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
