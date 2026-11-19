//go:build ignore

// 560. Subarray Sum Equals K (Medium)
// https://leetcode.com/problems/subarray-sum-equals-k/
// Run: go run 0560_subarray_sum_equals_k.go lc.go
package main

import (
	"fmt"
	"os"
)

func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	seen := map[int]int{0: 1} // prefix sum -> how many times seen
	for _, x := range nums {
		pre += x
		count += seen[pre-k]
		seen[pre]++
	}
	return count
}

func main() {
	assertEq(subarraySum([]int{1, 1, 1}, 2), 2)
	assertEq(subarraySum([]int{1, 2, 3}, 3), 2)
	assertEq(subarraySum([]int{1, -1, 0}, 0), 3)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
