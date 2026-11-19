//go:build ignore

// 410. Split Array Largest Sum (Hard)
// https://leetcode.com/problems/split-array-largest-sum/
// Run: go run 0410_split_array_largest_sum.go lc.go
package main

import (
	"fmt"
	"os"
)

func splitArray(nums []int, k int) int {
	parts := func(limit int) int { // pieces needed if no piece may exceed limit
		cnt, s := 1, 0
		for _, x := range nums {
			if s+x > limit {
				cnt, s = cnt+1, 0
			}
			s += x
		}
		return cnt
	}
	lo, hi := 0, 0
	for _, x := range nums {
		lo = max(lo, x)
		hi += x
	}
	for lo < hi {
		mid := lo + (hi-lo)/2
		if parts(mid) <= k {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	assertEq(splitArray([]int{7, 2, 5, 10, 8}, 2), 18)
	assertEq(splitArray([]int{1, 2, 3, 4, 5}, 2), 9)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
