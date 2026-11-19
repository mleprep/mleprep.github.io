//go:build ignore

// 2130. Maximum Twin Sum of a Linked List (Medium)
// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/
// Run: go run 2130_maximum_twin_sum_of_a_linked_list.go lc.go
package main

import (
	"fmt"
	"os"
)

func pairSum(head *ListNode) int {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	var prev *ListNode // reverse the second half
	for slow != nil {
		slow.Next, prev, slow = prev, slow, slow.Next
	}
	best := 0
	for ; prev != nil; head, prev = head.Next, prev.Next {
		best = max(best, head.Val+prev.Val)
	}
	return best
}

func main() {
	assertEq(pairSum(makeList(5, 4, 2, 1)), 6)
	assertEq(pairSum(makeList(4, 2, 2, 3)), 7)
	assertEq(pairSum(makeList(1, 100000)), 100001)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
