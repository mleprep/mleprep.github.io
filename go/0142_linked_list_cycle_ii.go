//go:build ignore

// 142. Linked List Cycle II (Medium)
// https://leetcode.com/problems/linked-list-cycle-ii/
// Run: go run 0142_linked_list_cycle_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast { // the cycle start is as far from head as from the meeting point
			p := head
			for p != slow {
				p, slow = p.Next, slow.Next
			}
			return p
		}
	}
	return nil
}

func main() {
	h := makeList(3, 2, 0, -4)
	h.Next.Next.Next.Next = h.Next
	assertEq(detectCycle(h) == h.Next, true)
	assertEq(detectCycle(makeList(1)) == nil, true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
