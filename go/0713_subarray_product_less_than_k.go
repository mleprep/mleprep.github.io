//go:build ignore

// 713. Subarray Product Less Than K (Medium)
// https://leetcode.com/problems/subarray-product-less-than-k/
// Run: go run 0713_subarray_product_less_than_k.go lc.go
package main

import (
	"fmt"
	"os"
)

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	prod, l, count := 1, 0, 0
	for r, x := range nums {
		prod *= x
		for prod >= k {
			prod /= nums[l]
			l++
		}
		count += r - l + 1 // subarrays ending at r
	}
	return count
}

func main() {
	assertEq(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100), 8)
	assertEq(numSubarrayProductLessThanK([]int{1, 2, 3}, 0), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
