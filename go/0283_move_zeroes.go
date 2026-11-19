//go:build ignore

// 283. Move Zeroes (Easy)
// https://leetcode.com/problems/move-zeroes/
// Run: go run 0283_move_zeroes.go lc.go
package main

import (
	"fmt"
	"os"
)

func moveZeroes(nums []int) {
	k := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	}
}

func main() {
	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)
	assertEq(a, []int{1, 3, 12, 0, 0})
	b := []int{0}
	moveZeroes(b)
	assertEq(b, []int{0})
	c := []int{1, 0, 0, 2}
	moveZeroes(c)
	assertEq(c, []int{1, 2, 0, 0})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
