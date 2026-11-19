//go:build ignore

// 1137. N-th Tribonacci Number (Easy)
// https://leetcode.com/problems/n-th-tribonacci-number/
// Run: go run 1137_n_th_tribonacci_number.go lc.go
package main

import (
	"fmt"
	"os"
)

func tribonacci(n int) int {
	a, b, c := 0, 1, 1
	for i := 0; i < n; i++ {
		a, b, c = b, c, a+b+c
	}
	return a
}

func main() {
	assertEq(tribonacci(4), 4)
	assertEq(tribonacci(25), 1389537)
	assertEq([]int{tribonacci(0), tribonacci(1), tribonacci(2), tribonacci(3)}, []int{0, 1, 1, 2})
	assertEq(tribonacci(37), 2082876103)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
