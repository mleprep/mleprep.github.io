//go:build ignore

// 1011. Capacity To Ship Packages Within D Days (Medium)
// https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/
// Run: go run 1011_capacity_to_ship_packages_within_d_days.go lc.go
package main

import (
	"fmt"
	"os"
)

func shipWithinDays(weights []int, days int) int {
	need := func(capacity int) int {
		d, load := 1, 0
		for _, w := range weights {
			if load+w > capacity {
				d, load = d+1, 0
			}
			load += w
		}
		return d
	}
	lo, hi := 0, 0
	for _, w := range weights {
		lo = max(lo, w)
		hi += w
	}
	for lo < hi { // smallest capacity that needs <= days
		mid := lo + (hi-lo)/2
		if need(mid) <= days {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	assertEq(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5), 15)
	assertEq(shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3), 6)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
