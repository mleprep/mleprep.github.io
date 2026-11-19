//go:build ignore

// 77. Combinations (Medium)
// https://leetcode.com/problems/combinations/
// Run: go run 0077_combinations.go lc.go
package main

import (
	"fmt"
	"os"
)

func combine(n int, k int) [][]int {
	res, path := [][]int{}, []int{}
	var bt func(start int)
	bt = func(start int) {
		if len(path) == k {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i <= n-(k-len(path))+1; i++ { // leave room for the remaining picks
			path = append(path, i)
			bt(i + 1)
			path = path[:len(path)-1]
		}
	}
	bt(1)
	return res
}

func main() {
	assertUnorderedDeep(combine(4, 2), [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}})
	assertUnorderedDeep(combine(1, 1), [][]int{{1}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
