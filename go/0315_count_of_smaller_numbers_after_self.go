//go:build ignore

// 315. Count of Smaller Numbers After Self (Hard)
// https://leetcode.com/problems/count-of-smaller-numbers-after-self/
// Run: go run 0315_count_of_smaller_numbers_after_self.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func countSmaller(nums []int) []int {
	vals := append([]int(nil), nums...)
	sort.Ints(vals)
	rank := map[int]int{} // value -> 1-based rank among distinct values
	for _, v := range vals {
		if _, ok := rank[v]; !ok {
			rank[v] = len(rank) + 1
		}
	}
	tree := make([]int, len(rank)+1) // Fenwick tree of counts by rank, filled right to left
	res := make([]int, len(nums))
	for k := len(nums) - 1; k >= 0; k-- {
		s := 0
		for i := rank[nums[k]] - 1; i > 0; i -= i & -i { // how many smaller values are to the right
			s += tree[i]
		}
		res[k] = s
		for i := rank[nums[k]]; i < len(tree); i += i & -i {
			tree[i]++
		}
	}
	return res
}

func main() {
	assertEq(countSmaller([]int{5, 2, 6, 1}), []int{2, 1, 1, 0})
	assertEq(countSmaller([]int{-1, -1}), []int{0, 0})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
