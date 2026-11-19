//go:build ignore

// 530. Minimum Absolute Difference in BST (Easy)
// https://leetcode.com/problems/minimum-absolute-difference-in-bst/
// Run: go run 0530_minimum_absolute_difference_in_bst.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func getMinimumDifference(root *TreeNode) int {
	best, prev, havePrev := math.MaxInt, 0, false
	var inorder func(*TreeNode) // in-order visits values sorted: compare neighbours
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if havePrev {
			best = min(best, node.Val-prev)
		}
		prev, havePrev = node.Val, true
		inorder(node.Right)
	}
	inorder(root)
	return best
}

func main() {
	assertEq(getMinimumDifference(makeTree(4, 2, 6, 1, 3)), 1)
	assertEq(getMinimumDifference(makeTree(1, 0, 48, nil, nil, 12, 49)), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
