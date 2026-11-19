//go:build ignore

// 739. Daily Temperatures (Medium)
// https://leetcode.com/problems/daily-temperatures/
// Run: go run 0739_daily_temperatures.go lc.go
package main

import (
	"fmt"
	"os"
)

func dailyTemperatures(temperatures []int) []int {
	res, st := make([]int, len(temperatures)), []int{} // stack of indices, temps decreasing
	for i, t := range temperatures {
		for len(st) > 0 && temperatures[st[len(st)-1]] < t {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			res[j] = i - j
		}
		st = append(st, i)
	}
	return res
}

func main() {
	assertEq(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}), []int{1, 1, 4, 2, 1, 1, 0, 0})
	assertEq(dailyTemperatures([]int{30, 40, 50, 60}), []int{1, 1, 1, 0})
	assertEq(dailyTemperatures([]int{60, 50, 40}), []int{0, 0, 0})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
