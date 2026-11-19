//go:build ignore

// 160. Intersection of Two Linked Lists (Easy)
// https://leetcode.com/problems/intersection-of-two-linked-lists/
// Run: go run 0160_intersection_of_two_linked_lists.go lc.go
package main

import (
	"fmt"
	"os"
)

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b { // both walk lenA + lenB nodes, so they meet at the junction (or at nil)
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}
		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}
	return a
}

func main() {
	common := makeList(8, 4, 5)
	A := makeList(4, 1)
	A.Next.Next = common
	B := makeList(5, 6, 1)
	B.Next.Next.Next = common
	assertEq(getIntersectionNode(A, B) == common, true)
	assertEq(getIntersectionNode(makeList(1), makeList(2)) == nil, true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
