//go:build ignore

// 1448. Count Good Nodes in Binary Tree (Medium)
// https://leetcode.com/problems/count-good-nodes-in-binary-tree/
// Run: go run 1448_count_good_nodes_in_binary_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func goodNodes(root *TreeNode) int {
	var dfs func(node *TreeNode, mx int) int
	dfs = func(node *TreeNode, mx int) int {
		if node == nil {
			return 0
		}
		good := 0
		if node.Val >= mx { // nothing bigger on the path from the root
			good = 1
		}
		mx = max(mx, node.Val)
		return good + dfs(node.Left, mx) + dfs(node.Right, mx)
	}
	return dfs(root, root.Val)
}

func main() {
	assertEq(goodNodes(makeTree(3, 1, 4, 3, nil, 1, 5)), 4)
	assertEq(goodNodes(makeTree(3, 3, nil, 4, 2)), 3)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
