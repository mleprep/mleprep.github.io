//go:build ignore

// 236. Lowest Common Ancestor of a Binary Tree (Medium)
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
// Run: go run 0236_lowest_common_ancestor_of_a_binary_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil { // one on each side
		return root
	}
	if l != nil {
		return l
	}
	return r
}

func main() {
	t := makeTree(3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4)
	assertEq(lowestCommonAncestor(t, findNode(t, 5), findNode(t, 1)).Val, 3)
	assertEq(lowestCommonAncestor(t, findNode(t, 5), findNode(t, 4)).Val, 5)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
