//go:build ignore

// 257. Binary Tree Paths (Easy)
// https://leetcode.com/problems/binary-tree-paths/
// Run: go run 0257_binary_tree_paths.go lc.go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		path += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			res = append(res, path)
			return
		}
		dfs(node.Left, path+"->")
		dfs(node.Right, path+"->")
	}
	dfs(root, "")
	return res
}

func main() {
	assertUnordered(binaryTreePaths(makeTree(1, 2, 3, nil, 5)), []string{"1->2->5", "1->3"})
	assertEq(binaryTreePaths(makeTree(1)), []string{"1"})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
