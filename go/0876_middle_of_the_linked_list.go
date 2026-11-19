//go:build ignore

// 876. Middle of the Linked List (Easy)
// https://leetcode.com/problems/middle-of-the-linked-list/
// Run: go run 0876_middle_of_the_linked_list.go lc.go
package main

import (
	"fmt"
	"os"
)

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

func main() {
	assertEq(middleNode(makeList(1, 2, 3, 4, 5)), []int{3, 4, 5})
	assertEq(middleNode(makeList(1, 2, 3, 4, 5, 6)), []int{4, 5, 6})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
