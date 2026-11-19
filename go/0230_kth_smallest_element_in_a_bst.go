//go:build ignore

// 230. Kth Smallest Element in a BST (Medium)
// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
// Run: go run 0230_kth_smallest_element_in_a_bst.go lc.go
package main

import (
	"fmt"
	"os"
)

func kthSmallest(root *TreeNode, k int) int {
	st, node := []*TreeNode{}, root // iterative in-order traversal
	for {
		for node != nil {
			st = append(st, node)
			node = node.Left
		}
		node = st[len(st)-1]
		st = st[:len(st)-1]
		if k--; k == 0 {
			return node.Val
		}
		node = node.Right
	}
}

func main() {
	assertEq(kthSmallest(makeTree(3, 1, 4, nil, 2), 1), 1)
	assertEq(kthSmallest(makeTree(5, 3, 6, 2, 4, nil, nil, 1), 3), 3)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
