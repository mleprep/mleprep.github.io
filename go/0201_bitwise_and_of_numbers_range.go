//go:build ignore

// 201. Bitwise AND of Numbers Range (Medium)
// https://leetcode.com/problems/bitwise-and-of-numbers-range/
// Run: go run 0201_bitwise_and_of_numbers_range.go lc.go
package main

import (
	"fmt"
	"os"
)

func rangeBitwiseAnd(left int, right int) int {
	shift := 0
	for left != right { // the answer is the common binary prefix
		left, right = left>>1, right>>1
		shift++
	}
	return left << shift
}

func main() {
	assertEq(rangeBitwiseAnd(5, 7), 4)
	assertEq(rangeBitwiseAnd(0, 0), 0)
	assertEq(rangeBitwiseAnd(1, 2147483647), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
