//go:build ignore

// 46. Permutations (Medium)
// https://leetcode.com/problems/permutations/
// Run: go run 0046_permutations.go lc.go
package main

import (
	"fmt"
	"os"
)

func permute(nums []int) [][]int {
	res, path, used := [][]int{}, []int{}, make([]bool, len(nums))
	var bt func()
	bt = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i, x := range nums {
			if !used[i] {
				used[i] = true
				path = append(path, x)
				bt()
				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}
	bt()
	return res
}

func main() {
	assertUnordered(permute([]int{1, 2, 3}), [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}})
	assertUnordered(permute([]int{0, 1}), [][]int{{0, 1}, {1, 0}})
	assertEq(permute([]int{1}), [][]int{{1}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
