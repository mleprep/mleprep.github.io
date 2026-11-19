//go:build ignore

// 70. Climbing Stairs (Easy)
// https://leetcode.com/problems/climbing-stairs/
// Run: go run 0070_climbing_stairs.go lc.go
package main

import (
	"fmt"
	"os"
)

func climbStairs(n int) int {
	a, b := 1, 1 // ways(0), ways(1)
	for i := 0; i < n-1; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	assertEq(climbStairs(2), 2)
	assertEq(climbStairs(5), 8)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
