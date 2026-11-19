//go:build ignore

// 86. Partition List (Medium)
// https://leetcode.com/problems/partition-list/
// Run: go run 0086_partition_list.go lc.go
package main

import (
	"fmt"
	"os"
)

func partition(head *ListNode, x int) *ListNode {
	small, big := &ListNode{}, &ListNode{}
	s, b := small, big
	for ; head != nil; head = head.Next {
		if head.Val < x {
			s.Next = head
			s = head
		} else {
			b.Next = head
			b = head
		}
	}
	b.Next = nil
	s.Next = big.Next
	return small.Next
}

func main() {
	assertEq(partition(makeList(1, 4, 3, 2, 5, 2), 3), []int{1, 2, 2, 4, 3, 5})
	assertEq(partition(makeList(2, 1), 2), []int{1, 2})
	assertEq(partition(makeList(), 0), []int{})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
