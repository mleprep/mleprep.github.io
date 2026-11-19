//go:build ignore

// 69. Sqrt(x) (Easy)
// https://leetcode.com/problems/sqrtx/
// Run: go run 0069_sqrtx.go lc.go
package main

import (
	"fmt"
	"os"
)

func mySqrt(x int) int {
	lo, hi := 0, x+1 // first m with m*m > x
	for lo < hi {
		m := lo + (hi-lo)/2
		if m*m > x {
			hi = m
		} else {
			lo = m + 1
		}
	}
	return lo - 1
}

func main() {
	assertEq(mySqrt(4), 2)
	assertEq(mySqrt(8), 2)
	assertEq(mySqrt(0), 0)
	assertEq(mySqrt(2147395599), 46339)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
