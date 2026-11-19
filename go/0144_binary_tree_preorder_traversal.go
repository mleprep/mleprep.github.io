//go:build ignore

// 144. Binary Tree Preorder Traversal (Easy)
// https://leetcode.com/problems/binary-tree-preorder-traversal/
// Run: go run 0144_binary_tree_preorder_traversal.go lc.go
package main

import (
	"fmt"
	"os"
)

func preorderTraversal(root *TreeNode) []int {
	out, st := []int{}, []*TreeNode{}
	if root != nil {
		st = append(st, root)
	}
	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		out = append(out, node.Val)
		if node.Right != nil { // push right first so left is visited first
			st = append(st, node.Right)
		}
		if node.Left != nil {
			st = append(st, node.Left)
		}
	}
	return out
}

func main() {
	assertEq(preorderTraversal(makeTree(1, nil, 2, 3)), []int{1, 2, 3})
	assertEq(preorderTraversal(makeTree(1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9)), []int{1, 2, 4, 5, 6, 7, 3, 8, 9})
	assertEq(preorderTraversal(makeTree()), []int{})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
