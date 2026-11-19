//go:build ignore

// 1431. Kids With the Greatest Number of Candies (Easy)
// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
// Run: go run 1431_kids_with_the_greatest_number_of_candies.go lc.go
package main

import (
	"fmt"
	"os"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	best := 0
	for _, c := range candies {
		best = max(best, c)
	}
	res := make([]bool, len(candies))
	for i, c := range candies {
		res[i] = c+extraCandies >= best
	}
	return res
}

func main() {
	assertEq(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3), []bool{true, true, true, false, true})
	assertEq(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1), []bool{true, false, false, false, false})
	assertEq(kidsWithCandies([]int{12, 1, 12}, 10), []bool{true, false, true})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
