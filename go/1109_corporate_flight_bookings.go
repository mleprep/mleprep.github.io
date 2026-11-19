//go:build ignore

// 1109. Corporate Flight Bookings (Medium)
// https://leetcode.com/problems/corporate-flight-bookings/
// Run: go run 1109_corporate_flight_bookings.go lc.go
package main

import (
	"fmt"
	"os"
)

func corpFlightBookings(bookings [][]int, n int) []int {
	d := make([]int, n+1) // difference array: +seats at first, -seats after last
	for _, b := range bookings {
		d[b[0]-1] += b[2]
		d[b[1]] -= b[2]
	}
	res := make([]int, n)
	run := 0
	for i := 0; i < n; i++ {
		run += d[i]
		res[i] = run
	}
	return res
}

func main() {
	assertEq(corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5), []int{10, 55, 45, 25, 25})
	assertEq(corpFlightBookings([][]int{{1, 2, 10}, {2, 2, 15}}, 2), []int{10, 25})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
