//go:build ignore

// 129. Sum Root to Leaf Numbers (Medium)
// https://leetcode.com/problems/sum-root-to-leaf-numbers/
// Run: go run 0129_sum_root_to_leaf_numbers.go lc.go
package main

import (
	"fmt"
	"os"
)

func sumNumbers(root *TreeNode) int {
	var dfs func(node *TreeNode, cur int) int
	dfs = func(node *TreeNode, cur int) int {
		if node == nil {
			return 0
		}
		cur = cur*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return cur
		}
		return dfs(node.Left, cur) + dfs(node.Right, cur)
	}
	return dfs(root, 0)
}

func main() {
	assertEq(sumNumbers(makeTree(1, 2, 3)), 25)
	assertEq(sumNumbers(makeTree(4, 9, 0, 5, 1)), 1026)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
