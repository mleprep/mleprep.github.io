//go:build ignore

// 109. Convert Sorted List to Binary Search Tree (Medium)
// https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
// Run: go run 0109_convert_sorted_list_to_binary_search_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func sortedListToBST(head *ListNode) *TreeNode {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	cur := head // consumed in order while building in-order
	var build func(lo, hi int) *TreeNode
	build = func(lo, hi int) *TreeNode {
		if lo > hi {
			return nil
		}
		mid := (lo + hi) / 2
		left := build(lo, mid-1)
		root := &TreeNode{Val: cur.Val, Left: left}
		cur = cur.Next
		root.Right = build(mid+1, hi)
		return root
	}
	return build(0, n-1)
}

func main() {
	// any height-balanced BST with the same in-order sequence is accepted
	var walk func(n *TreeNode) ([]int, int, bool) // in-order values, height, balanced
	walk = func(n *TreeNode) ([]int, int, bool) {
		if n == nil {
			return []int{}, 0, true
		}
		lv, lh, lok := walk(n.Left)
		rv, rh, rok := walk(n.Right)
		vals := append(append(lv, n.Val), rv...)
		return vals, 1 + max(lh, rh), lok && rok && lh-rh <= 1 && rh-lh <= 1
	}
	for _, vals := range [][]int{{-10, -3, 0, 5, 9}, {1, 3}, {0}, {1, 2, 3, 4, 5, 6, 7, 8}} {
		got, _, balanced := walk(sortedListToBST(makeList(vals...)))
		assertEq(balanced, true)
		assertEq(got, vals)
	}
	assertEq(sortedListToBST(nil) == nil, true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
