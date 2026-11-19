//go:build ignore

// 172. Factorial Trailing Zeroes (Medium)
// https://leetcode.com/problems/factorial-trailing-zeroes/
// Run: go run 0172_factorial_trailing_zeroes.go lc.go
package main

import (
	"fmt"
	"os"
)

func trailingZeroes(n int) int {
	count := 0
	for n > 0 { // count the factors of 5 (there are always more 2s)
		n /= 5
		count += n
	}
	return count
}

func main() {
	assertEq(trailingZeroes(3), 0)
	assertEq(trailingZeroes(5), 1)
	assertEq(trailingZeroes(100), 24)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
