//go:build ignore

// 238. Product of Array Except Self (Medium)
// https://leetcode.com/problems/product-of-array-except-self/
// Run: go run 0238_product_of_array_except_self.go lc.go
package main

import (
	"fmt"
	"os"
)

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] * nums[i-1] // product of everything to the left
	}
	right := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] *= right
		right *= nums[i]
	}
	return ans
}

func main() {
	assertEq(productExceptSelf([]int{1, 2, 3, 4}), []int{24, 12, 8, 6})
	assertEq(productExceptSelf([]int{-1, 1, 0, -3, 3}), []int{0, 0, 9, 0, 0})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
