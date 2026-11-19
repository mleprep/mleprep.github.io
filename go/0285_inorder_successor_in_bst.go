//go:build ignore

// 285. Inorder Successor in BST (Medium)
// https://leetcode.com/problems/inorder-successor-in-bst/
// Run: go run 0285_inorder_successor_in_bst.go lc.go
package main

import (
	"fmt"
	"os"
)

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var succ *TreeNode
	for root != nil {
		if p.Val < root.Val { // root is a candidate; look for a smaller one on the left
			succ, root = root, root.Left
		} else {
			root = root.Right
		}
	}
	return succ
}

func main() {
	t := makeTree(2, 1, 3)
	assertEq(inorderSuccessor(t, findNode(t, 1)).Val, 2)
	t2 := makeTree(5, 3, 6, 2, 4, nil, nil, 1)
	assertEq(inorderSuccessor(t2, findNode(t2, 6)) == nil, true)
	assertEq(inorderSuccessor(t2, findNode(t2, 4)).Val, 5)
	assertEq(inorderSuccessor(t2, findNode(t2, 1)).Val, 2)
	assertEq(inorderSuccessor(t2, findNode(t2, 3)).Val, 4)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
