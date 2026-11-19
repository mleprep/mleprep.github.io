//go:build ignore

// 94. Binary Tree Inorder Traversal (Easy)
// https://leetcode.com/problems/binary-tree-inorder-traversal/
// Run: go run 0094_binary_tree_inorder_traversal.go lc.go
package main

import (
	"fmt"
	"os"
)

func inorderTraversal(root *TreeNode) []int {
	res, st, node := []int{}, []*TreeNode{}, root
	for len(st) > 0 || node != nil {
		for node != nil { // go left as far as possible
			st = append(st, node)
			node = node.Left
		}
		node = st[len(st)-1]
		st = st[:len(st)-1]
		res = append(res, node.Val)
		node = node.Right
	}
	return res
}

func main() {
	assertEq(inorderTraversal(makeTree(1, nil, 2, 3)), []int{1, 3, 2})
	assertEq(inorderTraversal(makeTree(1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9)), []int{4, 2, 6, 5, 7, 1, 3, 9, 8})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
