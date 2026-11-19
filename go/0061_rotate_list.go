//go:build ignore

// 61. Rotate List (Medium)
// https://leetcode.com/problems/rotate-list/
// Run: go run 0061_rotate_list.go lc.go
package main

import (
	"fmt"
	"os"
)

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	n, tail := 1, head
	for tail.Next != nil {
		tail, n = tail.Next, n+1
	}
	tail.Next = head // close the ring, then cut it n - k%n nodes in
	for i := 0; i < n-k%n-1; i++ {
		head = head.Next
	}
	newHead := head.Next
	head.Next = nil
	return newHead
}

func main() {
	assertEq(rotateRight(makeList(1, 2, 3, 4, 5), 2), []int{4, 5, 1, 2, 3})
	assertEq(rotateRight(makeList(0, 1, 2), 4), []int{2, 0, 1})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
