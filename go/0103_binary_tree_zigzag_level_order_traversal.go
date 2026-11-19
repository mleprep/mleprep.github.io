//go:build ignore

// 103. Binary Tree Zigzag Level Order Traversal (Medium)
// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
// Run: go run 0103_binary_tree_zigzag_level_order_traversal.go lc.go
package main

import (
	"fmt"
	"os"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
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
		if len(res)%2 == 1 { // odd levels read right to left
			for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {
				level[i], level[j] = level[j], level[i]
			}
		}
		res = append(res, level)
	}
	return res
}

func main() {
	assertEq(zigzagLevelOrder(makeTree(3, 9, 20, nil, nil, 15, 7)), [][]int{{3}, {20, 9}, {15, 7}})
	assertEq(zigzagLevelOrder(makeTree(1)), [][]int{{1}})
	assertEq(zigzagLevelOrder(makeTree()), [][]int{})
	assertEq(zigzagLevelOrder(makeTree(1, 2, 3, 4, nil, nil, 5)), [][]int{{1}, {3, 2}, {4, 5}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
