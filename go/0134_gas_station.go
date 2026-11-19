//go:build ignore

// 134. Gas Station (Medium)
// https://leetcode.com/problems/gas-station/
// Run: go run 0134_gas_station.go lc.go
package main

import (
	"fmt"
	"os"
)

func canCompleteCircuit(gas []int, cost []int) int {
	total, tank, start := 0, 0, 0
	for i := range gas {
		total += gas[i] - cost[i]
		tank += gas[i] - cost[i]
		if tank < 0 { // cannot reach i+1 from start, nor from any station in between
			start, tank = i+1, 0
		}
	}
	if total < 0 {
		return -1
	}
	return start
}

func main() {
	assertEq(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}), 3)
	assertEq(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}), -1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
