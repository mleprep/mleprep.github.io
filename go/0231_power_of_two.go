//go:build ignore

// 231. Power of Two (Easy)
// https://leetcode.com/problems/power-of-two/
// Run: go run 0231_power_of_two.go lc.go
package main

import (
	"fmt"
	"os"
)

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0 // exactly one bit set
}

func main() {
	assertEq(isPowerOfTwo(16), true)
	assertEq(isPowerOfTwo(3), false)
	assertEq(isPowerOfTwo(-16), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
