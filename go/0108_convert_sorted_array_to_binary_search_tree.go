//go:build ignore

// 108. Convert Sorted Array to Binary Search Tree (Easy)
// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
// Run: go run 0108_convert_sorted_array_to_binary_search_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func sortedArrayToBST(nums []int) *TreeNode {
	var build func(lo, hi int) *TreeNode
	build = func(lo, hi int) *TreeNode {
		if lo > hi {
			return nil
		}
		mid := (lo + hi) / 2 // the middle element keeps both sides balanced
		return &TreeNode{Val: nums[mid], Left: build(lo, mid-1), Right: build(mid+1, hi)}
	}
	return build(0, len(nums)-1)
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
		got, _, balanced := walk(sortedArrayToBST(vals))
		assertEq(balanced, true)
		assertEq(got, vals)
	}
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
