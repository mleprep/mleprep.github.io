//go:build ignore

// 24. Swap Nodes in Pairs (Medium)
// https://leetcode.com/problems/swap-nodes-in-pairs/
// Run: go run 0024_swap_nodes_in_pairs.go lc.go
package main

import (
	"fmt"
	"os"
)

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for prev.Next != nil && prev.Next.Next != nil {
		a, b := prev.Next, prev.Next.Next
		prev.Next, a.Next, b.Next = b, b.Next, a
		prev = a
	}
	return dummy.Next
}

func main() {
	assertEq(swapPairs(makeList(1, 2, 3, 4)), []int{2, 1, 4, 3})
	assertEq(swapPairs(makeList(1, 2, 3)), []int{2, 1, 3})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
