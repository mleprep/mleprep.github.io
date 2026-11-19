//go:build ignore

// 112. Path Sum (Easy)
// https://leetcode.com/problems/path-sum/
// Run: go run 0112_path_sum.go lc.go
package main

import (
	"fmt"
	"os"
)

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	rest := targetSum - root.Val
	return hasPathSum(root.Left, rest) || hasPathSum(root.Right, rest)
}

func main() {
	assertEq(hasPathSum(makeTree(5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1), 22), true)
	assertEq(hasPathSum(makeTree(1, 2, 3), 5), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
