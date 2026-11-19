//go:build ignore

// 1049. Last Stone Weight II (Medium)
// https://leetcode.com/problems/last-stone-weight-ii/
// Run: go run 1049_last_stone_weight_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func lastStoneWeightII(stones []int) int {
	total := 0
	for _, x := range stones {
		total += x
	}
	half := total / 2
	reach := make([]bool, half+1) // reach[s]: some subset sums to s
	reach[0] = true
	for _, x := range stones {
		for s := half; s >= x; s-- {
			reach[s] = reach[s] || reach[s-x]
		}
	}
	best := half // split the stones into two piles as equal as possible
	for !reach[best] {
		best--
	}
	return total - 2*best
}

func main() {
	assertEq(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}), 1)
	assertEq(lastStoneWeightII([]int{31, 26, 33, 21, 40}), 5)
	assertEq(lastStoneWeightII([]int{1}), 1)
	assertEq(lastStoneWeightII([]int{1, 1}), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
