//go:build ignore

// 977. Squares of a Sorted Array (Easy)
// https://leetcode.com/problems/squares-of-a-sorted-array/
// Run: go run 0977_squares_of_a_sorted_array.go lc.go
package main

import (
	"fmt"
	"os"
)

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	l, r := 0, n-1
	for k := n - 1; k >= 0; k-- { // the biggest square is at one of the ends
		if nums[l]*nums[l] > nums[r]*nums[r] {
			res[k] = nums[l] * nums[l]
			l++
		} else {
			res[k] = nums[r] * nums[r]
			r--
		}
	}
	return res
}

func main() {
	assertEq(sortedSquares([]int{-4, -1, 0, 3, 10}), []int{0, 1, 9, 16, 100})
	assertEq(sortedSquares([]int{-7, -3, 2, 3, 11}), []int{4, 9, 9, 49, 121})
	assertEq(sortedSquares([]int{-5, -3, -2, -1}), []int{1, 4, 9, 25})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
