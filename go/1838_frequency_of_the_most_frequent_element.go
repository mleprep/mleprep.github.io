//go:build ignore

// 1838. Frequency of the Most Frequent Element (Medium)
// https://leetcode.com/problems/frequency-of-the-most-frequent-element/
// Run: go run 1838_frequency_of_the_most_frequent_element.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	l, s, best := 0, 0, 0
	for r, x := range nums {
		s += x
		for x*(r-l+1)-s > k { // cost to raise the window to x
			s -= nums[l]
			l++
		}
		best = max(best, r-l+1)
	}
	return best
}

func main() {
	assertEq(maxFrequency([]int{1, 2, 4}, 5), 3)
	assertEq(maxFrequency([]int{1, 4, 8, 13}, 5), 2)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
