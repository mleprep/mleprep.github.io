//go:build ignore

// 3133. Minimum Array End (Medium)
// https://leetcode.com/problems/minimum-array-end/
// Run: go run 3133_minimum_array_end.go lc.go
package main

import (
	"fmt"
	"os"
)

func minEnd(n int, x int) int64 {
	res, k := int64(x), int64(n-1)
	for bit := int64(1); k > 0; bit <<= 1 {
		if int64(x)&bit == 0 { // free position: place the next bit of n - 1
			res |= (k & 1) * bit
			k >>= 1
		}
	}
	return res
}

func main() {
	assertEq(minEnd(3, 4), 6)
	assertEq(minEnd(2, 7), 15)
	assertEq(minEnd(1, 5), 5)
	assertEq(minEnd(5, 1), 9)
	brute := func(n, x int) int { // n-th smallest number containing all bits of x
		for v, seen := x, 0; ; v++ {
			if v&x == x {
				seen++
				if seen == n {
					return v
				}
			}
		}
	}
	ok := true
	for n := 1; n < 25; n++ {
		for x := 1; x < 33; x++ {
			ok = ok && minEnd(n, x) == int64(brute(n, x))
		}
	}
	assertEq(ok, true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
