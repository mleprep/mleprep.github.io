//go:build ignore

// 23. Merge k Sorted Lists (Hard)
// https://leetcode.com/problems/merge-k-sorted-lists/
// Run: go run 0023_merge_k_sorted_lists.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"os"
)

type nodeHeap []*ListNode

func (h nodeHeap) Len() int           { return len(h) }
func (h nodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h nodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(x any)        { *h = append(*h, x.(*ListNode)) }
func (h *nodeHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &nodeHeap{}
	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}
	dummy := &ListNode{}
	tail := dummy
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode) // smallest head among the lists
		tail.Next = node
		tail = node
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return dummy.Next
}

func main() {
	assertEq(mergeKLists([]*ListNode{makeList(1, 4, 5), makeList(1, 3, 4), makeList(2, 6)}), []int{1, 1, 2, 3, 4, 4, 5, 6})
	assertEq(mergeKLists([]*ListNode{}) == nil, true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
