//go:build ignore

// 1899. Merge Triplets to Form Target Triplet (Medium)
// https://leetcode.com/problems/merge-triplets-to-form-target-triplet/
// Run: go run 1899_merge_triplets_to_form_target_triplet.go lc.go
package main

import (
	"fmt"
	"os"
)

func mergeTriplets(triplets [][]int, target []int) bool {
	good := [3]bool{}
	for _, t := range triplets {
		if t[0] > target[0] || t[1] > target[1] || t[2] > target[2] {
			continue // merging it would overshoot
		}
		for i := 0; i < 3; i++ {
			if t[i] == target[i] {
				good[i] = true
			}
		}
	}
	return good[0] && good[1] && good[2]
}

func main() {
	assertEq(mergeTriplets([][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, []int{2, 7, 5}), true)
	assertEq(mergeTriplets([][]int{{3, 4, 5}, {4, 5, 6}}, []int{3, 2, 5}), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
