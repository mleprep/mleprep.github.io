//go:build ignore

// 986. Interval List Intersections (Medium)
// https://leetcode.com/problems/interval-list-intersections/
// Run: go run 0986_interval_list_intersections.go lc.go
package main

import (
	"fmt"
	"os"
)

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := [][]int{}
	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {
		lo := max(firstList[i][0], secondList[j][0])
		hi := min(firstList[i][1], secondList[j][1])
		if lo <= hi {
			res = append(res, []int{lo, hi})
		}
		if firstList[i][1] < secondList[j][1] { // the one ending first cannot meet anything else
			i++
		} else {
			j++
		}
	}
	return res
}

func main() {
	A := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	B := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	assertEq(intervalIntersection(A, B), [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}})
	assertEq(intervalIntersection([][]int{{1, 3}, {5, 9}}, [][]int{}), [][]int{})
	assertEq(intervalIntersection([][]int{{1, 7}}, [][]int{{3, 10}}), [][]int{{3, 7}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
