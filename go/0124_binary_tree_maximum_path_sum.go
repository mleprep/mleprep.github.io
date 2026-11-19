//go:build ignore

// 124. Binary Tree Maximum Path Sum (Hard)
// https://leetcode.com/problems/binary-tree-maximum-path-sum/
// Run: go run 0124_binary_tree_maximum_path_sum.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func maxPathSum(root *TreeNode) int {
	best := math.MinInt
	var gain func(*TreeNode) int // best downward path starting at node
	gain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := max(gain(node.Left), 0), max(gain(node.Right), 0)
		best = max(best, node.Val+l+r)
		return node.Val + max(l, r)
	}
	gain(root)
	return best
}

func main() {
	assertEq(maxPathSum(makeTree(1, 2, 3)), 6)
	assertEq(maxPathSum(makeTree(-10, 9, 20, nil, nil, 15, 7)), 42)
	assertEq(maxPathSum(makeTree(-3)), -3)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
