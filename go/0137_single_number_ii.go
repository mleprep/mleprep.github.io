//go:build ignore

// 137. Single Number II (Medium)
// https://leetcode.com/problems/single-number-ii/
// Run: go run 0137_single_number_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func singleNumber(nums []int) int {
	res := 0
	for b := 0; b < 32; b++ { // each bit's count mod 3 is the single number's bit
		cnt := 0
		for _, x := range nums {
			cnt += (x >> b) & 1
		}
		if cnt%3 != 0 {
			res |= 1 << b
		}
	}
	return int(int32(res)) // bit 31 is the sign of a 32-bit int
}

func main() {
	assertEq(singleNumber([]int{2, 2, 3, 2}), 3)
	assertEq(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}), 99)
	assertEq(singleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}), -4)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
