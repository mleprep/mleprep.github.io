//go:build ignore

// 974. Subarray Sums Divisible by K (Medium)
// https://leetcode.com/problems/subarray-sums-divisible-by-k/
// Run: go run 0974_subarray_sums_divisible_by_k.go lc.go
package main

import (
	"fmt"
	"os"
)

func subarraysDivByK(nums []int, k int) int {
	cnt := make([]int, k)
	cnt[0] = 1
	pre, res := 0, 0
	for _, x := range nums {
		pre = ((pre+x)%k + k) % k // Go's % can be negative
		res += cnt[pre]
		cnt[pre]++
	}
	return res
}

func main() {
	assertEq(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5), 7)
	assertEq(subarraysDivByK([]int{5}, 9), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
