//go:build ignore

// 204. Count Primes (Medium)
// https://leetcode.com/problems/count-primes/
// Run: go run 0204_count_primes.go lc.go
package main

import (
	"fmt"
	"os"
)

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	composite := make([]bool, n) // sieve of Eratosthenes below n
	count := 0
	for p := 2; p < n; p++ {
		if composite[p] {
			continue
		}
		count++
		for j := p * p; j < n; j += p {
			composite[j] = true
		}
	}
	return count
}

func main() {
	assertEq(countPrimes(10), 4)
	assertEq(countPrimes(0), 0)
	assertEq(countPrimes(100), 25)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
