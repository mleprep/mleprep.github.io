//go:build ignore

// 228. Summary Ranges (Easy)
// https://leetcode.com/problems/summary-ranges/
// Run: go run 0228_summary_ranges.go lc.go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func summaryRanges(nums []int) []string {
	res := []string{}
	for i := 0; i < len(nums); {
		j := i
		for j+1 < len(nums) && nums[j+1] == nums[j]+1 {
			j++
		}
		if i == j {
			res = append(res, strconv.Itoa(nums[i]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[i], nums[j]))
		}
		i = j + 1
	}
	return res
}

func main() {
	assertEq(summaryRanges([]int{0, 1, 2, 4, 5, 7}), []string{"0->2", "4->5", "7"})
	assertEq(summaryRanges([]int{}), []string{})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
