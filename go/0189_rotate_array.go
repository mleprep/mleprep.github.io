//go:build ignore

// 189. Rotate Array (Medium)
// https://leetcode.com/problems/rotate-array/
// Run: go run 0189_rotate_array.go lc.go
package main

import (
	"fmt"
	"os"
)

func rotate(nums []int, k int) {
	rev := func(i, j int) {
		for ; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	n := len(nums)
	k %= n
	rev(0, n-1)
	rev(0, k-1)
	rev(k, n-1)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(a, 3)
	assertEq(a, []int{5, 6, 7, 1, 2, 3, 4})
	b := []int{-1, -100, 3, 99}
	rotate(b, 2)
	assertEq(b, []int{3, 99, -1, -100})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
