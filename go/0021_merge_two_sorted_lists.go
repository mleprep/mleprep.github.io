//go:build ignore

// 21. Merge Two Sorted Lists (Easy)
// https://leetcode.com/problems/merge-two-sorted-lists/
// Run: go run 0021_merge_two_sorted_lists.go lc.go
package main

import (
	"fmt"
	"os"
)

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next, list1 = list1, list1.Next
		} else {
			tail.Next, list2 = list2, list2.Next
		}
		tail = tail.Next
	}
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}
	return dummy.Next
}

func main() {
	assertEq(mergeTwoLists(makeList(1, 2, 4), makeList(1, 3, 4)), []int{1, 1, 2, 3, 4, 4})
	assertEq(mergeTwoLists(makeList(), makeList()), []int{})
	assertEq(mergeTwoLists(makeList(), makeList(0)), []int{0})
	assertEq(mergeTwoLists(makeList(5), makeList(1, 2, 4)), []int{1, 2, 4, 5})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
