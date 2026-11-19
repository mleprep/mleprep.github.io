//go:build ignore

// 525. Contiguous Array (Medium)
// https://leetcode.com/problems/contiguous-array/
// Run: go run 0525_contiguous_array.go lc.go
package main

import (
	"fmt"
	"os"
)

func findMaxLength(nums []int) int {
	first := map[int]int{0: -1} // balance -> first index it was seen
	pre, best := 0, 0
	for i, x := range nums {
		if x == 1 {
			pre++
		} else {
			pre--
		}
		if j, ok := first[pre]; ok {
			best = max(best, i-j)
		} else {
			first[pre] = i
		}
	}
	return best
}

func main() {
	assertEq(findMaxLength([]int{0, 1}), 2)
	assertEq(findMaxLength([]int{0, 1, 1, 0, 1, 1}), 4)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
