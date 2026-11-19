//go:build ignore

// 202. Happy Number (Easy)
// https://leetcode.com/problems/happy-number/
// Run: go run 0202_happy_number.go lc.go
package main

import (
	"fmt"
	"os"
)

func isHappy(n int) bool {
	next := func(x int) int { // sum of the squares of the digits
		s := 0
		for ; x > 0; x /= 10 {
			s += (x % 10) * (x % 10)
		}
		return s
	}
	slow, fast := n, next(n) // Floyd: either reach 1 or loop
	for fast != 1 && slow != fast {
		slow, fast = next(slow), next(next(fast))
	}
	return fast == 1
}

func main() {
	assertEq(isHappy(19), true)
	assertEq(isHappy(2), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
