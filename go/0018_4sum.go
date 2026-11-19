//go:build ignore

// 18. 4Sum (Medium)
// https://leetcode.com/problems/4sum/
// Run: go run 0018_4sum.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ksum func(start, k, target int) [][]int
	ksum = func(start, k, target int) [][]int {
		res := [][]int{}
		if k == 2 {
			for l, r := start, len(nums)-1; l < r; {
				s := nums[l] + nums[r]
				if s < target || (l > start && nums[l] == nums[l-1]) {
					l++
				} else if s > target {
					r--
				} else {
					res = append(res, []int{nums[l], nums[r]})
					l, r = l+1, r-1
				}
			}
			return res
		}
		for i := start; i <= len(nums)-k; i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			for _, sub := range ksum(i+1, k-1, target-nums[i]) {
				res = append(res, append([]int{nums[i]}, sub...))
			}
		}
		return res
	}
	return ksum(0, 4, target)
}

func main() {
	assertUnorderedDeep(fourSum([]int{1, 0, -1, 0, -2, 2}, 0), [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}})
	assertUnorderedDeep(fourSum([]int{2, 2, 2, 2, 2}, 8), [][]int{{2, 2, 2, 2}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
