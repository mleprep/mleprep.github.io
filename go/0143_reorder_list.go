//go:build ignore

// 143. Reorder List (Medium)
// https://leetcode.com/problems/reorder-list/
// Run: go run 0143_reorder_list.go lc.go
package main

import (
	"fmt"
	"os"
)

func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil { // find the middle
		slow, fast = slow.Next, fast.Next.Next
	}
	var prev *ListNode // reverse the second half
	cur := slow.Next
	slow.Next = nil
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	for a, b := head, prev; b != nil; { // weave the halves together
		an, bn := a.Next, b.Next
		a.Next, b.Next = b, an
		a, b = an, bn
	}
}

func main() {
	h := makeList(1, 2, 3, 4, 5)
	reorderList(h)
	assertEq(h, []int{1, 5, 2, 4, 3})
	h2 := makeList(1, 2, 3, 4)
	reorderList(h2)
	assertEq(h2, []int{1, 4, 2, 3})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
