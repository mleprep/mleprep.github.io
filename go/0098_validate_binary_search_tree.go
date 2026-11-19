//go:build ignore

// 98. Validate Binary Search Tree (Medium)
// https://leetcode.com/problems/validate-binary-search-tree/
// Run: go run 0098_validate_binary_search_tree.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func isValidBST(root *TreeNode) bool {
	var ok func(node *TreeNode, lo, hi int) bool // every value must lie in (lo, hi)
	ok = func(node *TreeNode, lo, hi int) bool {
		if node == nil {
			return true
		}
		if node.Val <= lo || node.Val >= hi {
			return false
		}
		return ok(node.Left, lo, node.Val) && ok(node.Right, node.Val, hi)
	}
	return ok(root, math.MinInt, math.MaxInt)
}

func main() {
	assertEq(isValidBST(makeTree(2, 1, 3)), true)
	assertEq(isValidBST(makeTree(5, 1, 6, nil, nil, 3, 7)), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
