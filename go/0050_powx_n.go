//go:build ignore

// 50. Pow(x, n) (Medium)
// https://leetcode.com/problems/powx-n/
// Run: go run 0050_powx_n.go lc.go
package main

import (
	"fmt"
	"os"
)

func myPow(x float64, n int) float64 {
	if n < 0 {
		x, n = 1/x, -n
	}
	res := 1.0
	for ; n > 0; n >>= 1 { // binary exponentiation
		if n&1 == 1 {
			res *= x
		}
		x *= x
	}
	return res
}

func main() {
	assertEq(myPow(2.0, 10), 1024.0)
	assertEq(myPow(2.1, 3), 9.261)
	assertEq(myPow(2.0, -2), 0.25)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
