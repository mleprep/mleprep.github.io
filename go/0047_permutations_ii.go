//go:build ignore

// 47. Permutations II (Medium)
// https://leetcode.com/problems/permutations-ii/
// Run: go run 0047_permutations_ii.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res, path, used := [][]int{}, []int{}, make([]bool, len(nums))
	var bt func()
	bt = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := range nums {
			// equal values are used left to right only, so each arrangement appears once
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			bt()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	bt()
	return res
}

func main() {
	assertUnordered(permuteUnique([]int{1, 1, 2}), [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}})
	assertUnordered(permuteUnique([]int{1, 2, 3}), [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}})
	assertEq(permuteUnique([]int{2, 2}), [][]int{{2, 2}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
