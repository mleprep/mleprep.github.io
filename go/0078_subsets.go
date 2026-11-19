//go:build ignore

// 78. Subsets (Medium)
// https://leetcode.com/problems/subsets/
// Run: go run 0078_subsets.go lc.go
package main

import (
	"fmt"
	"os"
)

func subsets(nums []int) [][]int {
	res, path := [][]int{}, []int{}
	var bt func(start int)
	bt = func(start int) {
		res = append(res, append([]int{}, path...)) // every node of the tree is a subset
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			bt(i + 1)
			path = path[:len(path)-1]
		}
	}
	bt(0)
	return res
}

func main() {
	assertUnorderedDeep(subsets([]int{1, 2, 3}), [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}})
	assertUnorderedDeep(subsets([]int{0}), [][]int{{}, {0}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
