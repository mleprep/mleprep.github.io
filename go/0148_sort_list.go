//go:build ignore

// 148. Sort List (Medium)
// https://leetcode.com/problems/sort-list/
// Run: go run 0148_sort_list.go lc.go
package main

import (
	"fmt"
	"os"
)

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil // split, sort both halves, merge
	a, b := sortList(head), sortList(mid)
	dummy := &ListNode{}
	tail := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next, a = a, a.Next
		} else {
			tail.Next, b = b, b.Next
		}
		tail = tail.Next
	}
	if a != nil {
		tail.Next = a
	} else {
		tail.Next = b
	}
	return dummy.Next
}

func main() {
	assertEq(sortList(makeList(4, 2, 1, 3)), []int{1, 2, 3, 4})
	assertEq(sortList(makeList(-1, 5, 3, 4, 0)), []int{-1, 0, 3, 4, 5})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
