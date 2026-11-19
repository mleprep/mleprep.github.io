//go:build ignore

// 328. Odd Even Linked List (Medium)
// https://leetcode.com/problems/odd-even-linked-list/
// Run: go run 0328_odd_even_linked_list.go lc.go
package main

import (
	"fmt"
	"os"
)

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd, even := head, head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

func main() {
	assertEq(oddEvenList(makeList(1, 2, 3, 4, 5)), []int{1, 3, 5, 2, 4})
	assertEq(oddEvenList(makeList(2, 1, 3, 5, 6, 4, 7)), []int{2, 3, 6, 7, 1, 5, 4})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
