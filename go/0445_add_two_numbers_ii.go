//go:build ignore

// 445. Add Two Numbers II (Medium)
// https://leetcode.com/problems/add-two-numbers-ii/
// Run: go run 0445_add_two_numbers_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var a, b []int // digits, most significant first
	for ; l1 != nil; l1 = l1.Next {
		a = append(a, l1.Val)
	}
	for ; l2 != nil; l2 = l2.Next {
		b = append(b, l2.Val)
	}
	var head *ListNode
	carry := 0
	for len(a) > 0 || len(b) > 0 || carry > 0 {
		s := carry
		if len(a) > 0 {
			s += a[len(a)-1]
			a = a[:len(a)-1]
		}
		if len(b) > 0 {
			s += b[len(b)-1]
			b = b[:len(b)-1]
		}
		carry = s / 10
		head = &ListNode{Val: s % 10, Next: head} // build from the least significant end
	}
	return head
}

func main() {
	assertEq(addTwoNumbers(makeList(7, 2, 4, 3), makeList(5, 6, 4)), []int{7, 8, 0, 7})
	assertEq(addTwoNumbers(makeList(5), makeList(5)), []int{1, 0})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
