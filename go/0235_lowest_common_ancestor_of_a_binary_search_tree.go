//go:build ignore

// 235. Lowest Common Ancestor of a Binary Search Tree (Medium)
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
// Run: go run 0235_lowest_common_ancestor_of_a_binary_search_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		switch {
		case p.Val < root.Val && q.Val < root.Val:
			root = root.Left
		case p.Val > root.Val && q.Val > root.Val:
			root = root.Right
		default: // p and q split here (or one of them is root)
			return root
		}
	}
	return nil
}

func main() {
	t := makeTree(6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5)
	assertEq(lowestCommonAncestor(t, findNode(t, 2), findNode(t, 8)).Val, 6)
	assertEq(lowestCommonAncestor(t, findNode(t, 2), findNode(t, 4)).Val, 2)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
