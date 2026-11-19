//go:build ignore

// 1. Two Sum (Easy)
// https://leetcode.com/problems/two-sum/
// Run: go run 0001_two_sum.go lc.go
package main

import (
	"fmt"
	"os"
)

func twoSum(nums []int, target int) []int {
	seen := map[int]int{} // value -> index
	for i, x := range nums {
		if j, ok := seen[target-x]; ok {
			return []int{j, i}
		}
		seen[x] = i
	}
	return nil
}

func main() {
	assertUnordered(twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1})
	assertUnordered(twoSum([]int{3, 2, 4}, 6), []int{1, 2})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
