//go:build ignore

// 92. Reverse Linked List II (Medium)
// https://leetcode.com/problems/reverse-linked-list-ii/
// Run: go run 0092_reverse_linked_list_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	cur := prev.Next
	for i := 0; i < right-left; i++ { // move the node after cur to the front of the range
		nxt := cur.Next
		cur.Next = nxt.Next
		nxt.Next = prev.Next
		prev.Next = nxt
	}
	return dummy.Next
}

func main() {
	assertEq(reverseBetween(makeList(1, 2, 3, 4, 5), 2, 4), []int{1, 4, 3, 2, 5})
	assertEq(reverseBetween(makeList(5), 1, 1), []int{5})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
