//go:build ignore

// 1325. Delete Leaves With a Given Value (Medium)
// https://leetcode.com/problems/delete-leaves-with-a-given-value/
// Run: go run 1325_delete_leaves_with_a_given_value.go lc.go
package main

import (
	"fmt"
	"os"
)

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = removeLeafNodes(root.Left, target) // children first: a parent may become a leaf
	root.Right = removeLeafNodes(root.Right, target)
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}

func main() {
	assertEq(removeLeafNodes(makeTree(1, 2, 3, 2, nil, 2, 4), 2), []any{1, nil, 3, nil, 4})
	assertEq(removeLeafNodes(makeTree(1, 3, 3, 3, 2), 3), []any{1, 3, nil, nil, 2})
	assertEq(removeLeafNodes(makeTree(1, 2, nil, 2, nil, 2), 2), []any{1})
	assertEq(removeLeafNodes(makeTree(1, 1, 1), 1) == nil, true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
