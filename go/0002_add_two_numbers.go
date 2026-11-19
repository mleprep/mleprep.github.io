//go:build ignore

// 2. Add Two Numbers (Medium)
// https://leetcode.com/problems/add-two-numbers/
// Run: go run 0002_add_two_numbers.go lc.go
package main

import (
	"fmt"
	"os"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail, carry := dummy, 0
	for l1 != nil || l2 != nil || carry > 0 {
		s := carry
		if l1 != nil {
			s, l1 = s+l1.Val, l1.Next
		}
		if l2 != nil {
			s, l2 = s+l2.Val, l2.Next
		}
		carry = s / 10
		tail.Next = &ListNode{Val: s % 10}
		tail = tail.Next
	}
	return dummy.Next
}

func main() {
	assertEq(addTwoNumbers(makeList(2, 4, 3), makeList(5, 6, 4)), []int{7, 0, 8})
	assertEq(addTwoNumbers(makeList(9, 9, 9), makeList(1)), []int{0, 0, 0, 1})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
