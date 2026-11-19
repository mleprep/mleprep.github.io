//go:build ignore

// 102. Binary Tree Level Order Traversal (Medium)
// https://leetcode.com/problems/binary-tree-level-order-traversal/
// Run: go run 0102_binary_tree_level_order_traversal.go lc.go
package main

import (
	"fmt"
	"os"
)

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		level := []int{}
		for n := len(q); n > 0; n-- {
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

func main() {
	assertEq(levelOrder(makeTree(3, 9, 20, nil, nil, 15, 7)), [][]int{{3}, {9, 20}, {15, 7}})
	assertEq(levelOrder(nil), [][]int{})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
