//go:build ignore

// 523. Continuous Subarray Sum (Medium)
// https://leetcode.com/problems/continuous-subarray-sum/
// Run: go run 0523_continuous_subarray_sum.go lc.go
package main

import (
	"fmt"
	"os"
)

func checkSubarraySum(nums []int, k int) bool {
	first := map[int]int{0: -1} // prefix sum mod k -> first index
	pre := 0
	for i, x := range nums {
		pre = (pre + x) % k
		if j, ok := first[pre]; ok {
			if i-j >= 2 {
				return true
			}
		} else {
			first[pre] = i
		}
	}
	return false
}

func main() {
	assertEq(checkSubarraySum([]int{23, 2, 4, 6, 7}, 6), true)
	assertEq(checkSubarraySum([]int{23, 2, 6, 4, 7}, 13), false)
	assertEq(checkSubarraySum([]int{5, 0, 0, 0}, 3), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
