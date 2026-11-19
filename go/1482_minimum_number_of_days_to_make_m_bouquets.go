//go:build ignore

// 1482. Minimum Number of Days to Make m Bouquets (Medium)
// https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/
// Run: go run 1482_minimum_number_of_days_to_make_m_bouquets.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	ok := func(day int) bool {
		made, run := 0, 0
		for _, b := range bloomDay {
			if b <= day {
				run++
			} else {
				run = 0
			}
			if run == k {
				made, run = made+1, 0
			}
		}
		return made >= m
	}
	lo, hi := math.MaxInt, 0
	for _, b := range bloomDay {
		lo, hi = min(lo, b), max(hi, b)
	}
	for lo < hi {
		mid := lo + (hi-lo)/2
		if ok(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	assertEq(minDays([]int{1, 10, 3, 10, 2}, 3, 1), 3)
	assertEq(minDays([]int{1, 10, 3, 10, 2}, 3, 2), -1)
	assertEq(minDays([]int{7, 7, 7, 7, 12, 7, 7}, 2, 3), 12)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
