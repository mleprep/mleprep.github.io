//go:build ignore

// 29. Divide Two Integers (Medium)
// https://leetcode.com/problems/divide-two-integers/
// Run: go run 0029_divide_two_integers.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32 // the only overflowing case
	}
	neg := (dividend < 0) != (divisor < 0)
	a, b, q := dividend, divisor, 0
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for a >= b { // subtract the largest b * 2^k that fits
		k := 0
		for a >= b<<(k+1) {
			k++
		}
		a -= b << k
		q += 1 << k
	}
	if neg {
		return -q
	}
	return q
}

func main() {
	assertEq(divide(10, 3), 3)
	assertEq(divide(7, -3), -2)
	assertEq(divide(-2147483648, -1), 2147483647)
	assertEq(divide(-2147483648, 2), -1073741824)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
