//go:build ignore

// 105. Construct Binary Tree from Preorder and Inorder Traversal (Medium)
// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// Run: go run 0105_construct_binary_tree_from_preorder_and_inorder_traversal.go lc.go
package main

import (
	"fmt"
	"os"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	pos := map[int]int{}
	for i, v := range inorder {
		pos[v] = i
	}
	next := 0 // next preorder value = next subtree root
	var build func(lo, hi int) *TreeNode
	build = func(lo, hi int) *TreeNode { // inorder range [lo, hi)
		if lo >= hi {
			return nil
		}
		root := &TreeNode{Val: preorder[next]}
		next++
		m := pos[root.Val]
		root.Left = build(lo, m) // the left subtree comes first in preorder
		root.Right = build(m+1, hi)
		return root
	}
	return build(0, len(inorder))
}

func main() {
	assertEq(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}), []any{3, 9, 20, nil, nil, 15, 7})
	assertEq(buildTree([]int{-1}, []int{-1}), []any{-1})
	assertEq(buildTree([]int{1, 2, 3}, []int{3, 2, 1}), []any{1, 2, nil, 3})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
