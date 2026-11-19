//go:build ignore

// 40. Combination Sum II (Medium)
// https://leetcode.com/problems/combination-sum-ii/
// Run: go run 0040_combination_sum_ii.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res, path := [][]int{}, []int{}
	var bt func(start, rest int)
	bt = func(start, rest int) {
		if rest == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i < len(candidates) && candidates[i] <= rest; i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			bt(i+1, rest-candidates[i])
			path = path[:len(path)-1]
		}
	}
	bt(0, target)
	return res
}

func main() {
	assertUnorderedDeep(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8), [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}})
	assertUnorderedDeep(combinationSum2([]int{2, 5, 2, 1, 2}, 5), [][]int{{1, 2, 2}, {5}})
	assertEq(combinationSum2([]int{1}, 2), [][]int{})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
