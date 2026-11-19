//go:build ignore

// 1094. Car Pooling (Medium)
// https://leetcode.com/problems/car-pooling/
// Run: go run 1094_car_pooling.go lc.go
package main

import (
	"fmt"
	"os"
)

func carPooling(trips [][]int, capacity int) bool {
	d := make([]int, 1001) // passenger change at each location
	for _, t := range trips {
		d[t[1]] += t[0]
		d[t[2]] -= t[0] // they leave at t[2], before anyone boards there
	}
	load := 0
	for _, x := range d {
		load += x
		if load > capacity {
			return false
		}
	}
	return true
}

func main() {
	assertEq(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4), false)
	assertEq(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5), true)
	assertEq(carPooling([][]int{{2, 1, 5}, {3, 5, 7}}, 3), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
