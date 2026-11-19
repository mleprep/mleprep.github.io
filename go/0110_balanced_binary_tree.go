//go:build ignore

// 110. Balanced Binary Tree (Easy)
// https://leetcode.com/problems/balanced-binary-tree/
// Run: go run 0110_balanced_binary_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func isBalanced(root *TreeNode) bool {
	var h func(*TreeNode) int // height, or -1 if unbalanced
	h = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := h(node.Left), h(node.Right)
		if l < 0 || r < 0 || l-r > 1 || r-l > 1 {
			return -1
		}
		return 1 + max(l, r)
	}
	return h(root) >= 0
}

func main() {
	assertEq(isBalanced(makeTree(3, 9, 20, nil, nil, 15, 7)), true)
	assertEq(isBalanced(makeTree(1, 2, 2, 3, 3, nil, nil, 4, 4)), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
