//go:build ignore

// 1207. Unique Number of Occurrences (Easy)
// https://leetcode.com/problems/unique-number-of-occurrences/
// Run: go run 1207_unique_number_of_occurrences.go lc.go
package main

import (
	"fmt"
	"os"
)

func uniqueOccurrences(arr []int) bool {
	cnt := map[int]int{}
	for _, x := range arr {
		cnt[x]++
	}
	seen := map[int]bool{}
	for _, c := range cnt {
		if seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
}

func main() {
	assertEq(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}), true)
	assertEq(uniqueOccurrences([]int{1, 2}), false)
	assertEq(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
