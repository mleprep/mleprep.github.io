//go:build ignore

// 25. Reverse Nodes in k-Group (Hard)
// https://leetcode.com/problems/reverse-nodes-in-k-group/
// Run: go run 0025_reverse_nodes_in_k_group.go lc.go
package main

import (
	"fmt"
	"os"
)

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	groupPrev := dummy
	for {
		kth := groupPrev
		for i := 0; i < k; i++ {
			kth = kth.Next
			if kth == nil { // fewer than k nodes left: keep them as they are
				return dummy.Next
			}
		}
		groupNext := kth.Next
		prev, cur := groupNext, groupPrev.Next // reverse [groupPrev.Next .. kth]
		for cur != groupNext {
			cur.Next, prev, cur = prev, cur, cur.Next
		}
		first := groupPrev.Next
		groupPrev.Next = kth
		groupPrev = first
	}
}

func main() {
	assertEq(reverseKGroup(makeList(1, 2, 3, 4, 5), 2), []int{2, 1, 4, 3, 5})
	assertEq(reverseKGroup(makeList(1, 2, 3, 4, 5), 3), []int{3, 2, 1, 4, 5})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
