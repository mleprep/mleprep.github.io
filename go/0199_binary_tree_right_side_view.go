//go:build ignore

// 199. Binary Tree Right Side View (Medium)
// https://leetcode.com/problems/binary-tree-right-side-view/
// Run: go run 0199_binary_tree_right_side_view.go lc.go
package main

import (
	"fmt"
	"os"
)

func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		res = append(res, q[len(q)-1].Val) // last node of each level
		for n := len(q); n > 0; n-- {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return res
}

func main() {
	assertEq(rightSideView(makeTree(1, 2, 3, nil, 5, nil, 4)), []int{1, 3, 4})
	assertEq(rightSideView(makeTree(1, 2, 3, 4)), []int{1, 3, 4})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
