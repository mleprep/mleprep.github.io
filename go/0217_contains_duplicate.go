//go:build ignore

// 217. Contains Duplicate (Easy)
// https://leetcode.com/problems/contains-duplicate/
// Run: go run 0217_contains_duplicate.go lc.go
package main

import (
	"fmt"
	"os"
)

func containsDuplicate(nums []int) bool {
	seen := map[int]bool{}
	for _, x := range nums {
		if seen[x] {
			return true
		}
		seen[x] = true
	}
	return false
}

func main() {
	assertEq(containsDuplicate([]int{1, 2, 3, 1}), true)
	assertEq(containsDuplicate([]int{1, 2, 3, 4}), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
