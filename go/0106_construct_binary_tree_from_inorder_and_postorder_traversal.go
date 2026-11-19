//go:build ignore

// 106. Construct Binary Tree from Inorder and Postorder Traversal (Medium)
// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
// Run: go run 0106_construct_binary_tree_from_inorder_and_postorder_traversal.go lc.go
package main

import (
	"fmt"
	"os"
)

func buildTree(inorder []int, postorder []int) *TreeNode {
	pos := map[int]int{}
	for i, v := range inorder {
		pos[v] = i
	}
	next := len(postorder) - 1 // postorder read backwards: root, right subtree, left subtree
	var build func(lo, hi int) *TreeNode
	build = func(lo, hi int) *TreeNode {
		if lo >= hi {
			return nil
		}
		root := &TreeNode{Val: postorder[next]}
		next--
		m := pos[root.Val]
		root.Right = build(m+1, hi)
		root.Left = build(lo, m)
		return root
	}
	return build(0, len(inorder))
}

func main() {
	assertEq(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}), []any{3, 9, 20, nil, nil, 15, 7})
	assertEq(buildTree([]int{-1}, []int{-1}), []any{-1})
	assertEq(buildTree([]int{2, 1}, []int{2, 1}), []any{1, 2})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
