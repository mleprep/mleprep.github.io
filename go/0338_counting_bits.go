//go:build ignore

// 338. Counting Bits (Easy)
// https://leetcode.com/problems/counting-bits/
// Run: go run 0338_counting_bits.go lc.go
package main

import (
	"fmt"
	"os"
)

func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i>>1] + i&1 // i's bits = (i/2)'s bits + the last bit
	}
	return bits
}

func main() {
	assertEq(countBits(5), []int{0, 1, 1, 2, 1, 2})
	assertEq(countBits(0), []int{0})
	assertEq(countBits(2), []int{0, 1, 1})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
