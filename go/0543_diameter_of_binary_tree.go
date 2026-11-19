//go:build ignore

// 543. Diameter of Binary Tree (Easy)
// https://leetcode.com/problems/diameter-of-binary-tree/
// Run: go run 0543_diameter_of_binary_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func diameterOfBinaryTree(root *TreeNode) int {
	best := 0
	var height func(*TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := height(node.Left), height(node.Right)
		best = max(best, l+r) // longest path through node
		return 1 + max(l, r)
	}
	height(root)
	return best
}

func main() {
	assertEq(diameterOfBinaryTree(makeTree(1, 2, 3, 4, 5)), 3)
	assertEq(diameterOfBinaryTree(makeTree(1, 2)), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
