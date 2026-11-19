//go:build ignore

// 2095. Delete the Middle Node of a Linked List (Medium)
// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/
// Run: go run 2095_delete_the_middle_node_of_a_linked_list.go lc.go
package main

import (
	"fmt"
	"os"
)

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	slow, fast := head, head.Next.Next // slow stops just before the middle
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func main() {
	assertEq(deleteMiddle(makeList(1, 3, 4, 7, 1, 2, 6)), []int{1, 3, 4, 1, 2, 6})
	assertEq(deleteMiddle(makeList(1, 2, 3, 4)), []int{1, 2, 4})
	assertEq(deleteMiddle(makeList(2, 1)), []int{2})
	assertEq(deleteMiddle(makeList(1)), []int{})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
