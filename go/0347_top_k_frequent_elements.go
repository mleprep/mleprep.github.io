//go:build ignore

// 347. Top K Frequent Elements (Medium)
// https://leetcode.com/problems/top-k-frequent-elements/
// Run: go run 0347_top_k_frequent_elements.go lc.go
package main

import (
	"fmt"
	"os"
)

func topKFrequent(nums []int, k int) []int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	bucket := make([][]int, len(nums)+1) // bucket[f] = values seen f times
	for x, f := range cnt {
		bucket[f] = append(bucket[f], x)
	}
	res := []int{}
	for f := len(bucket) - 1; f > 0; f-- {
		for _, x := range bucket[f] {
			res = append(res, x)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}

func main() {
	assertUnordered(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2), []int{1, 2})
	assertEq(topKFrequent([]int{1}, 1), []int{1})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
