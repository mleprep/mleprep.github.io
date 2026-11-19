//go:build ignore

// 448. Find All Numbers Disappeared in an Array (Easy)
// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
// Run: go run 0448_find_all_numbers_disappeared_in_an_array.go lc.go
package main

import (
	"fmt"
	"os"
)

func findDisappearedNumbers(nums []int) []int {
	for _, v := range nums {
		i := abs(v) - 1
		nums[i] = -abs(nums[i]) // mark value i+1 as seen
	}
	res := []int{}
	for i, v := range nums {
		if v > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	assertEq(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}), []int{5, 6})
	assertEq(findDisappearedNumbers([]int{1, 1}), []int{2})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
