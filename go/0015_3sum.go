//go:build ignore

// 15. 3Sum (Medium)
// https://leetcode.com/problems/3sum/
// Run: go run 0015_3sum.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // same first number, same triples
		}
		for l, r := i+1, len(nums)-1; l < r; {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < 0:
				l++
			case s > 0:
				r--
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				r--
			}
		}
	}
	return res
}

func main() {
	assertUnorderedDeep(threeSum([]int{-1, 0, 1, 2, -1, -4}), [][]int{{-1, -1, 2}, {-1, 0, 1}})
	assertUnorderedDeep(threeSum([]int{0, 0, 0}), [][]int{{0, 0, 0}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
