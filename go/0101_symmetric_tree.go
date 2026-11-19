//go:build ignore

// 101. Symmetric Tree (Easy)
// https://leetcode.com/problems/symmetric-tree/
// Run: go run 0101_symmetric_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func isSymmetric(root *TreeNode) bool {
	var mirror func(a, b *TreeNode) bool
	mirror = func(a, b *TreeNode) bool {
		if a == nil || b == nil {
			return a == b
		}
		return a.Val == b.Val && mirror(a.Left, b.Right) && mirror(a.Right, b.Left)
	}
	return mirror(root.Left, root.Right)
}

func main() {
	assertEq(isSymmetric(makeTree(1, 2, 2, 3, 4, 4, 3)), true)
	assertEq(isSymmetric(makeTree(1, 2, 2, nil, 3, nil, 3)), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
