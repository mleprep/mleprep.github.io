//go:build ignore

// 82. Remove Duplicates from Sorted List II (Medium)
// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
// Run: go run 0082_remove_duplicates_from_sorted_list_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for prev.Next != nil {
		cur := prev.Next
		if cur.Next != nil && cur.Next.Val == cur.Val {
			for cur.Next != nil && cur.Next.Val == cur.Val {
				cur = cur.Next
			}
			prev.Next = cur.Next // drop the whole run
		} else {
			prev = cur
		}
	}
	return dummy.Next
}

func main() {
	assertEq(deleteDuplicates(makeList(1, 2, 3, 3, 4, 4, 5)), []int{1, 2, 5})
	assertEq(deleteDuplicates(makeList(1, 1, 1, 2, 3)), []int{2, 3})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
