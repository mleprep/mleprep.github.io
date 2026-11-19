//go:build ignore

// 191. Number of 1 Bits (Easy)
// https://leetcode.com/problems/number-of-1-bits/
// Run: go run 0191_number_of_1_bits.go lc.go
package main

import (
	"fmt"
	"os"
)

func hammingWeight(n int) int {
	count := 0
	for ; n != 0; n &= n - 1 { // n & (n-1) clears the lowest set bit
		count++
	}
	return count
}

func main() {
	assertEq(hammingWeight(11), 3)
	assertEq(hammingWeight(2147483645), 30)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
