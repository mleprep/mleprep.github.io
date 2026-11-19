//go:build ignore

// 90. Subsets II (Medium)
// https://leetcode.com/problems/subsets-ii/
// Run: go run 0090_subsets_ii.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res, path := [][]int{}, []int{}
	var bt func(start int)
	bt = func(start int) {
		res = append(res, append([]int{}, path...))
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] { // same value at the same depth: skip
				continue
			}
			path = append(path, nums[i])
			bt(i + 1)
			path = path[:len(path)-1]
		}
	}
	bt(0)
	return res
}

func main() {
	assertUnorderedDeep(subsetsWithDup([]int{1, 2, 2}), [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}})
	assertUnorderedDeep(subsetsWithDup([]int{0}), [][]int{{}, {0}})
	assertUnorderedDeep(subsetsWithDup([]int{4, 4, 1, 4}), [][]int{{}, {1}, {1, 4}, {1, 4, 4}, {1, 4, 4, 4}, {4}, {4, 4}, {4, 4, 4}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
