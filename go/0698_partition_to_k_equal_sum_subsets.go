//go:build ignore

// 698. Partition to K Equal Sum Subsets (Medium)
// https://leetcode.com/problems/partition-to-k-equal-sum-subsets/
// Run: go run 0698_partition_to_k_equal_sum_subsets.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	total := 0
	for _, x := range nums {
		total += x
	}
	if total%k != 0 {
		return false
	}
	target := total / k
	sort.Sort(sort.Reverse(sort.IntSlice(nums))) // big numbers first fail fast
	if nums[0] > target {
		return false
	}
	buckets := make([]int, k)
	var bt func(i int) bool
	bt = func(i int) bool {
		if i == len(nums) {
			return true
		}
		for b := 0; b < k; b++ {
			if buckets[b]+nums[i] <= target {
				buckets[b] += nums[i]
				if bt(i + 1) {
					return true
				}
				buckets[b] -= nums[i]
			}
			if buckets[b] == 0 { // empty buckets are interchangeable
				break
			}
		}
		return false
	}
	return bt(0)
}

func main() {
	assertEq(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4), true)
	assertEq(canPartitionKSubsets([]int{1, 2, 3, 4}, 3), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
