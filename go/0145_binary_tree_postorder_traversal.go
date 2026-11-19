//go:build ignore

// 145. Binary Tree Postorder Traversal (Easy)
// https://leetcode.com/problems/binary-tree-postorder-traversal/
// Run: go run 0145_binary_tree_postorder_traversal.go lc.go
package main

import (
	"fmt"
	"os"
)

func postorderTraversal(root *TreeNode) []int {
	out, st := []int{}, []*TreeNode{}
	if root != nil {
		st = append(st, root)
	}
	for len(st) > 0 { // root-right-left, then reversed = left-right-root
		node := st[len(st)-1]
		st = st[:len(st)-1]
		out = append(out, node.Val)
		if node.Left != nil {
			st = append(st, node.Left)
		}
		if node.Right != nil {
			st = append(st, node.Right)
		}
	}
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}

func main() {
	assertEq(postorderTraversal(makeTree(1, nil, 2, 3)), []int{3, 2, 1})
	assertEq(postorderTraversal(makeTree(1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9)), []int{4, 6, 7, 5, 2, 9, 8, 3, 1})
	assertEq(postorderTraversal(makeTree()), []int{})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
