//go:build ignore

// 875. Koko Eating Bananas (Medium)
// https://leetcode.com/problems/koko-eating-bananas/
// Run: go run 0875_koko_eating_bananas.go lc.go
package main

import (
	"fmt"
	"os"
)

func minEatingSpeed(piles []int, h int) int {
	lo, hi := 1, 0
	for _, p := range piles {
		hi = max(hi, p)
	}
	for lo < hi { // smallest speed that finishes in h hours
		k := lo + (hi-lo)/2
		hours := 0
		for _, p := range piles {
			hours += (p + k - 1) / k
		}
		if hours <= h {
			hi = k
		} else {
			lo = k + 1
		}
	}
	return lo
}

func main() {
	assertEq(minEatingSpeed([]int{3, 6, 7, 11}, 8), 4)
	assertEq(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5), 30)
	assertEq(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6), 23)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
