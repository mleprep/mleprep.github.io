//go:build ignore

// 104. Maximum Depth of Binary Tree (Easy)
// https://leetcode.com/problems/maximum-depth-of-binary-tree/
// Run: go run 0104_maximum_depth_of_binary_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func main() {
	assertEq(maxDepth(makeTree(3, 9, 20, nil, nil, 15, 7)), 3)
	assertEq(maxDepth(nil), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
