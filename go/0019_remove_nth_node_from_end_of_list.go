//go:build ignore

// 19. Remove Nth Node From End of List (Medium)
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// Run: go run 0019_remove_nth_node_from_end_of_list.go lc.go
package main

import (
	"fmt"
	"os"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for i := 0; i < n; i++ { // fast runs n nodes ahead
		fast = fast.Next
	}
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func main() {
	assertEq(removeNthFromEnd(makeList(1, 2, 3, 4, 5), 2), []int{1, 2, 3, 5})
	assertEq(removeNthFromEnd(makeList(1), 1), []int{})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
