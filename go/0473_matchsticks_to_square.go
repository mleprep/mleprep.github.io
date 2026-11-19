//go:build ignore

// 473. Matchsticks to Square (Medium)
// https://leetcode.com/problems/matchsticks-to-square/
// Run: go run 0473_matchsticks_to_square.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func makesquare(matchsticks []int) bool {
	total := 0
	for _, x := range matchsticks {
		total += x
	}
	if total%4 != 0 {
		return false
	}
	side := total / 4
	sticks := append([]int{}, matchsticks...)
	sort.Sort(sort.Reverse(sort.IntSlice(sticks)))
	if sticks[0] > side {
		return false
	}
	var sides [4]int
	var place func(i int) bool
	place = func(i int) bool {
		if i == len(sticks) { // every side <= side and the total is 4*side, so all are equal
			return true
		}
		for k := 0; k < 4; k++ {
			if sides[k]+sticks[i] <= side {
				sides[k] += sticks[i]
				if place(i + 1) {
					return true
				}
				sides[k] -= sticks[i]
			}
			if sides[k] == 0 {
				break
			}
		}
		return false
	}
	return place(0)
}

func main() {
	assertEq(makesquare([]int{1, 1, 2, 2, 2}), true)
	assertEq(makesquare([]int{3, 3, 3, 3, 4}), false)
	assertEq(makesquare([]int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}), true)
	assertEq(makesquare([]int{10, 6, 5, 5, 5, 3, 3, 3, 2, 2, 2, 2}), true)
	assertEq(makesquare([]int{1, 1, 1}), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
