//go:build ignore

// 790. Domino and Tromino Tiling (Medium)
// https://leetcode.com/problems/domino-and-tromino-tiling/
// Run: go run 0790_domino_and_tromino_tiling.go lc.go
package main

import (
	"fmt"
	"os"
)

func numTilings(n int) int {
	const mod = 1_000_000_007
	if n < 3 {
		return []int{1, 1, 2}[n]
	}
	a, b, c := 1, 1, 2 // f(n) = 2 f(n-1) + f(n-3)
	for i := 0; i < n-2; i++ {
		a, b, c = b, c, (2*c+a)%mod
	}
	return c
}

func main() {
	got := []int{}
	for n := 1; n <= 6; n++ {
		got = append(got, numTilings(n))
	}
	assertEq(got, []int{1, 2, 5, 11, 24, 53})
	// cross-check against the two-state (full / jagged) DP for larger n
	f, p := []int{1, 1}, []int{0, 1} // p(n): a 2 x n board plus one protruding cell
	for n := 2; n <= 1000; n++ {
		f = append(f, (f[n-1]+f[n-2]+2*p[n-2])%1_000_000_007)
		p = append(p, (p[n-1]+f[n-1])%1_000_000_007)
	}
	assertEq(numTilings(1000), f[1000])
	assertEq(numTilings(30), f[30])
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
