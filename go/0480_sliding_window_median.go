//go:build ignore

// 480. Sliding Window Median (Hard)
// https://leetcode.com/problems/sliding-window-median/
// Run: go run 0480_sliding_window_median.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func medianSlidingWindow(nums []int, k int) []float64 {
	win := append([]int{}, nums[:k]...) // the window, kept sorted
	sort.Ints(win)
	res := []float64{}
	for i := k; ; i++ {
		res = append(res, (float64(win[k/2])+float64(win[(k-1)/2]))/2)
		if i == len(nums) {
			return res
		}
		j := sort.SearchInts(win, nums[i-k]) // remove the outgoing value
		win = append(win[:j], win[j+1:]...)
		j = sort.SearchInts(win, nums[i]) // insert the incoming one
		win = append(win, 0)
		copy(win[j+1:], win[j:])
		win[j] = nums[i]
	}
}

func main() {
	assertEq(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3), []float64{1, -1, -1, 3, 5, 6})
	assertEq(medianSlidingWindow([]int{1, 2, 3, 4, 2, 3, 1, 4, 2}, 3), []float64{2, 3, 3, 3, 2, 3, 2})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
