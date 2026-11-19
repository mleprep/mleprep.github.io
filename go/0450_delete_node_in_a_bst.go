//go:build ignore

// 450. Delete Node in a BST (Medium)
// https://leetcode.com/problems/delete-node-in-a-bst/
// Run: go run 0450_delete_node_in_a_bst.go lc.go
package main

import (
	"fmt"
	"os"
)

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	switch {
	case key < root.Val:
		root.Left = deleteNode(root.Left, key)
	case key > root.Val:
		root.Right = deleteNode(root.Right, key)
	default:
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		succ := root.Right // two children: take the in-order successor's value
		for succ.Left != nil {
			succ = succ.Left
		}
		root.Val = succ.Val
		root.Right = deleteNode(root.Right, succ.Val)
	}
	return root
}

func main() {
	assertEq(deleteNode(makeTree(5, 3, 6, 2, 4, nil, 7), 3), []any{5, 4, 6, 2, nil, nil, 7})
	assertEq(deleteNode(makeTree(5, 3, 6, 2, 4, nil, 7), 0), []any{5, 3, 6, 2, 4, nil, 7})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
