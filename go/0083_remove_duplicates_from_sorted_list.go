//go:build ignore

// 83. Remove Duplicates from Sorted List (Easy)
// https://leetcode.com/problems/remove-duplicates-from-sorted-list/
// Run: go run 0083_remove_duplicates_from_sorted_list.go lc.go
package main

import (
	"fmt"
	"os"
)

func deleteDuplicates(head *ListNode) *ListNode {
	for cur := head; cur != nil && cur.Next != nil; {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func main() {
	assertEq(deleteDuplicates(makeList(1, 1, 2, 3, 3)), []int{1, 2, 3})
	assertEq(deleteDuplicates(makeList()), []int{})
	assertEq(deleteDuplicates(makeList(1, 1, 1)), []int{1})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
