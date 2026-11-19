//go:build ignore

// 39. Combination Sum (Medium)
// https://leetcode.com/problems/combination-sum/
// Run: go run 0039_combination_sum.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res, path := [][]int{}, []int{}
	var bt func(start, rest int)
	bt = func(start, rest int) {
		if rest == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i < len(candidates) && candidates[i] <= rest; i++ {
			path = append(path, candidates[i])
			bt(i, rest-candidates[i]) // i: the same number may be reused
			path = path[:len(path)-1]
		}
	}
	bt(0, target)
	return res
}

func main() {
	assertUnorderedDeep(combinationSum([]int{2, 3, 6, 7}, 7), [][]int{{2, 2, 3}, {7}})
	assertUnorderedDeep(combinationSum([]int{2, 3, 5}, 8), [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
