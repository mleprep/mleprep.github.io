//go:build ignore

// 287. Find the Duplicate Number (Medium)
// https://leetcode.com/problems/find-the-duplicate-number/
// Run: go run 0287_find_the_duplicate_number.go lc.go
package main

import (
	"fmt"
	"os"
)

func findDuplicate(nums []int) int {
	slow, fast := 0, 0 // i -> nums[i] is a linked list whose cycle starts at the duplicate
	for {
		slow, fast = nums[slow], nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	p := 0
	for p != slow {
		p, slow = nums[p], nums[slow]
	}
	return p
}

func main() {
	assertEq(findDuplicate([]int{1, 3, 4, 2, 2}), 2)
	assertEq(findDuplicate([]int{3, 1, 3, 4, 2}), 3)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
