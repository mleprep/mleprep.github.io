//go:build ignore

// 141. Linked List Cycle (Easy)
// https://leetcode.com/problems/linked-list-cycle/
// Run: go run 0141_linked_list_cycle.go lc.go
package main

import (
	"fmt"
	"os"
)

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	h := makeList(3, 2, 0, -4)
	h.Next.Next.Next.Next = h.Next // tail -> node 1
	assertEq(hasCycle(h), true)
	assertEq(hasCycle(makeList(1)), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
