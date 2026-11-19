//go:build ignore

// 371. Sum of Two Integers (Medium)
// https://leetcode.com/problems/sum-of-two-integers/
// Run: go run 0371_sum_of_two_integers.go lc.go
package main

import (
	"fmt"
	"os"
)

func getSum(a int, b int) int {
	for b != 0 { // sum without carry, then the carry shifted left; ints wrap, so negatives work
		a, b = a^b, (a&b)<<1
	}
	return a
}

func main() {
	assertEq(getSum(1, 2), 3)
	assertEq(getSum(2, 3), 5)
	assertEq(getSum(-1, 1), 0)
	assertEq(getSum(-12, -8), -20)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
