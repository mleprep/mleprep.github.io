//go:build ignore

// 1318. Minimum Flips to Make a OR b Equal to c (Medium)
// https://leetcode.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/
// Run: go run 1318_minimum_flips_to_make_a_or_b_equal_to_c.go lc.go
package main

import (
	"fmt"
	"math/bits"
	"os"
)

func minFlips(a int, b int, c int) int {
	flips := 0
	for a > 0 || b > 0 || c > 0 {
		x, y, z := a&1, b&1, c&1
		if z == 1 {
			flips += (x | y) ^ 1 // need at least one 1
		} else {
			flips += x + y // both must be 0
		}
		a, b, c = a>>1, b>>1, c>>1
	}
	return flips
}

func main() {
	assertEq(minFlips(2, 6, 5), 3)
	assertEq(minFlips(4, 2, 7), 1)
	assertEq(minFlips(1, 2, 3), 0)
	assertEq(minFlips(8, 3, 5), 3)
	ok := true
	for a := 0; a < 16; a++ {
		for b := 0; b < 16; b++ {
			for c := 0; c < 16; c++ {
				ok = ok && minFlips(a, b, c) == bits.OnesCount(uint((a|b)^c))+bits.OnesCount(uint(a&b&^c))
			}
		}
	}
	assertEq(ok, true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
