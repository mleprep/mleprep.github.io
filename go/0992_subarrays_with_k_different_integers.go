//go:build ignore

// 992. Subarrays with K Different Integers (Hard)
// https://leetcode.com/problems/subarrays-with-k-different-integers/
// Run: go run 0992_subarrays_with_k_different_integers.go lc.go
package main

import (
	"fmt"
	"os"
)

func subarraysWithKDistinct(nums []int, k int) int {
	atMost := func(k int) int { // subarrays with at most k distinct values
		cnt := map[int]int{}
		l, res := 0, 0
		for r, x := range nums {
			cnt[x]++
			for len(cnt) > k {
				cnt[nums[l]]--
				if cnt[nums[l]] == 0 {
					delete(cnt, nums[l])
				}
				l++
			}
			res += r - l + 1
		}
		return res
	}
	return atMost(k) - atMost(k-1)
}

func main() {
	assertEq(subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2), 7)
	assertEq(subarraysWithKDistinct([]int{1, 2, 1, 3, 4}, 3), 3)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
