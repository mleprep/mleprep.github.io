//go:build ignore

// 2807. Insert Greatest Common Divisors in Linked List (Medium)
// https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list/
// Run: go run 2807_insert_greatest_common_divisors_in_linked_list.go lc.go
package main

import (
	"fmt"
	"os"
)

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	for cur := head; cur != nil && cur.Next != nil; cur = cur.Next.Next {
		cur.Next = &ListNode{Val: gcd(cur.Val, cur.Next.Val), Next: cur.Next}
	}
	return head
}

func main() {
	assertEq(insertGreatestCommonDivisors(makeList(18, 6, 10, 3)), []int{18, 6, 6, 2, 10, 1, 3})
	assertEq(insertGreatestCommonDivisors(makeList(7)), []int{7})
	assertEq(insertGreatestCommonDivisors(makeList(4, 4)), []int{4, 4, 4})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
