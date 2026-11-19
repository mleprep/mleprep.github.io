//go:build ignore

// 700. Search in a Binary Search Tree (Easy)
// https://leetcode.com/problems/search-in-a-binary-search-tree/
// Run: go run 0700_search_in_a_binary_search_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

func main() {
	assertEq(searchBST(makeTree(4, 2, 7, 1, 3), 2), []any{2, 1, 3})
	assertEq(searchBST(makeTree(4, 2, 7, 1, 3), 5) == nil, true)
	assertEq(searchBST(makeTree(4, 2, 7, 1, 3), 7), []any{7})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
