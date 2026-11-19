//go:build ignore

// 368. Largest Divisible Subset (Medium)
// https://leetcode.com/problems/largest-divisible-subset/
// Run: go run 0368_largest_divisible_subset.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	size, parent := make([]int, n), make([]int, n) // chain length ending at i, previous index
	best := 0
	for i := 0; i < n; i++ {
		size[i], parent[i] = 1, -1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && size[j]+1 > size[i] {
				size[i], parent[i] = size[j]+1, j
			}
		}
		if size[i] > size[best] {
			best = i
		}
	}
	res := []int{}
	for i := best; i != -1; i = parent[i] {
		res = append([]int{nums[i]}, res...)
	}
	return res
}

func main() {
	got := largestDivisibleSubset([]int{1, 2, 3})
	assertEq(len(got) == 2 && got[0] == 1 && (got[1] == 2 || got[1] == 3), true)
	assertEq(largestDivisibleSubset([]int{1, 2, 4, 8}), []int{1, 2, 4, 8})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
