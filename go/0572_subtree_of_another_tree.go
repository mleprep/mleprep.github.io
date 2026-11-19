//go:build ignore

// 572. Subtree of Another Tree (Easy)
// https://leetcode.com/problems/subtree-of-another-tree/
// Run: go run 0572_subtree_of_another_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return same(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func same(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Val == b.Val && same(a.Left, b.Left) && same(a.Right, b.Right)
}

func main() {
	assertEq(isSubtree(makeTree(3, 4, 5, 1, 2), makeTree(4, 1, 2)), true)
	assertEq(isSubtree(makeTree(3, 4, 5, 1, 2, nil, nil, nil, nil, 0), makeTree(4, 1, 2)), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
