//go:build ignore

// 111. Minimum Depth of Binary Tree (Easy)
// https://leetcode.com/problems/minimum-depth-of-binary-tree/
// Run: go run 0111_minimum_depth_of_binary_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q, depth := []*TreeNode{root}, 1
	for { // BFS: the first leaf found is the shallowest
		next := []*TreeNode{}
		for _, n := range q {
			if n.Left == nil && n.Right == nil {
				return depth
			}
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		q, depth = next, depth+1
	}
}

func main() {
	assertEq(minDepth(makeTree(3, 9, 20, nil, nil, 15, 7)), 2)
	assertEq(minDepth(makeTree(2, nil, 3, nil, 4, nil, 5, nil, 6)), 5)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
