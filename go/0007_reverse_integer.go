//go:build ignore

// 7. Reverse Integer (Medium)
// https://leetcode.com/problems/reverse-integer/
// Run: go run 0007_reverse_integer.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func reverse(x int) int {
	res := 0
	for ; x != 0; x /= 10 { // Go's % keeps the sign, so negatives just work
		res = res*10 + x%10
	}
	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}
	return res
}

func main() {
	assertEq(reverse(123), 321)
	assertEq(reverse(-123), -321)
	assertEq(reverse(120), 21)
	assertEq(reverse(1534236469), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
