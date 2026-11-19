//go:build ignore

// 226. Invert Binary Tree (Easy)
// https://leetcode.com/problems/invert-binary-tree/
// Run: go run 0226_invert_binary_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}

func main() {
	assertEq(invertTree(makeTree(4, 2, 7, 1, 3, 6, 9)), []any{4, 7, 2, 9, 6, 3, 1})
	assertEq(invertTree(makeTree()), []any{})
	assertEq(invertTree(makeTree(2, 1, 3)), []any{2, 3, 1})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
