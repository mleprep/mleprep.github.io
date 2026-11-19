//go:build ignore

// 16. 3Sum Closest (Medium)
// https://leetcode.com/problems/3sum-closest/
// Run: go run 0016_3sum_closest.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	best := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		for l, r := i+1, len(nums)-1; l < r; {
			s := nums[i] + nums[l] + nums[r]
			if abs(s-target) < abs(best-target) {
				best = s
			}
			switch {
			case s < target:
				l++
			case s > target:
				r--
			default:
				return s
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	assertEq(threeSumClosest([]int{-1, 2, 1, -4}, 1), 2)
	assertEq(threeSumClosest([]int{0, 0, 0}, 1), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
