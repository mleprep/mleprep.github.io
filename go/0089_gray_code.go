//go:build ignore

// 89. Gray Code (Medium)
// https://leetcode.com/problems/gray-code/
// Run: go run 0089_gray_code.go lc.go
package main

import (
	"fmt"
	"math/bits"
	"os"
)

func grayCode(n int) []int {
	res := make([]int, 1<<n)
	for i := range res {
		res[i] = i ^ i>>1 // reflected binary code
	}
	return res
}

func main() {
	// any valid Gray code: starts at 0, visits all 2^n values once, neighbours (incl. wrap-around) differ by one bit
	for n := 1; n <= 4; n++ {
		g := grayCode(n)
		ok := len(g) == 1<<n && g[0] == 0
		seen := map[int]bool{}
		for i := 0; ok && i < len(g); i++ {
			ok = g[i] >= 0 && g[i] < 1<<n && !seen[g[i]] && bits.OnesCount(uint(g[i]^g[(i+1)%len(g)])) == 1
			seen[g[i]] = true
		}
		assertEq(ok, true)
	}
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
