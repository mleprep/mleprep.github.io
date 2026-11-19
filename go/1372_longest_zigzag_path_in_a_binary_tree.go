//go:build ignore

// 1372. Longest ZigZag Path in a Binary Tree (Medium)
// https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/
// Run: go run 1372_longest_zigzag_path_in_a_binary_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func longestZigZag(root *TreeNode) int {
	best := 0
	var dfs func(*TreeNode) (int, int) // zigzag length (in edges) starting with a left / right step
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return -1, -1
		}
		_, lr := dfs(node.Left)
		rl, _ := dfs(node.Right)
		left, right := lr+1, rl+1
		best = max(best, left, right)
		return left, right
	}
	dfs(root)
	return best
}

func main() {
	assertEq(longestZigZag(makeTree(1, nil, 1, 1, 1, nil, nil, 1, 1, nil, 1, nil, nil, nil, 1)), 3)
	assertEq(longestZigZag(makeTree(1, 1, 1, nil, 1, nil, nil, 1, 1, nil, 1)), 4)
	assertEq(longestZigZag(makeTree(1)), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
