//go:build ignore

// 206. Reverse Linked List (Easy)
// https://leetcode.com/problems/reverse-linked-list/
// Run: go run 0206_reverse_linked_list.go lc.go
package main

import (
	"fmt"
	"os"
)

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for cur := head; cur != nil; {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}

func main() {
	assertEq(reverseList(makeList(1, 2, 3, 4, 5)), []int{5, 4, 3, 2, 1})
	assertEq(reverseList(nil) == nil, true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
