//go:build ignore

// 239. Sliding Window Maximum (Hard)
// https://leetcode.com/problems/sliding-window-maximum/
// Run: go run 0239_sliding_window_maximum.go lc.go
package main

import (
	"fmt"
	"os"
)

func maxSlidingWindow(nums []int, k int) []int {
	dq, res := []int{}, []int{} // indices, values decreasing
	for i, x := range nums {
		for len(dq) > 0 && nums[dq[len(dq)-1]] <= x {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
		if dq[0] <= i-k {
			dq = dq[1:]
		}
		if i >= k-1 {
			res = append(res, nums[dq[0]])
		}
	}
	return res
}

func main() {
	assertEq(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3), []int{3, 3, 5, 5, 6, 7})
	assertEq(maxSlidingWindow([]int{1}, 1), []int{1})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
