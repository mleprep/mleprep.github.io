//go:build ignore

// 169. Majority Element (Easy)
// https://leetcode.com/problems/majority-element/
// Run: go run 0169_majority_element.go lc.go
package main

import (
	"fmt"
	"os"
)

func majorityElement(nums []int) int {
	cand, cnt := 0, 0 // Boyer-Moore voting
	for _, x := range nums {
		if cnt == 0 {
			cand = x
		}
		if x == cand {
			cnt++
		} else {
			cnt--
		}
	}
	return cand
}

func main() {
	assertEq(majorityElement([]int{3, 2, 3}), 3)
	assertEq(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}), 2)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
