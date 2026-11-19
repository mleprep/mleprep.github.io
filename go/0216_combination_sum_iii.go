//go:build ignore

// 216. Combination Sum III (Medium)
// https://leetcode.com/problems/combination-sum-iii/
// Run: go run 0216_combination_sum_iii.go lc.go
package main

import (
	"fmt"
	"os"
)

func combinationSum3(k int, n int) [][]int {
	res, path := [][]int{}, []int{}
	var bt func(start, rest int)
	bt = func(start, rest int) {
		if len(path) == k {
			if rest == 0 {
				res = append(res, append([]int{}, path...))
			}
			return
		}
		for d := start; d <= 9 && d <= rest; d++ {
			path = append(path, d)
			bt(d+1, rest-d)
			path = path[:len(path)-1]
		}
	}
	bt(1, n)
	return res
}

func main() {
	assertUnorderedDeep(combinationSum3(3, 7), [][]int{{1, 2, 4}})
	assertUnorderedDeep(combinationSum3(3, 9), [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}})
	assertEq(combinationSum3(4, 1), [][]int{})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
