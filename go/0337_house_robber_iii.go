//go:build ignore

// 337. House Robber III (Medium)
// https://leetcode.com/problems/house-robber-iii/
// Run: go run 0337_house_robber_iii.go lc.go
package main

import (
	"fmt"
	"os"
)

func rob(root *TreeNode) int {
	var dfs func(*TreeNode) (int, int) // (best if node is robbed, best if it is skipped)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		lr, ls := dfs(node.Left)
		rr, rs := dfs(node.Right)
		return node.Val + ls + rs, max(lr, ls) + max(rr, rs)
	}
	robbed, skipped := dfs(root)
	return max(robbed, skipped)
}

func main() {
	assertEq(rob(makeTree(3, 2, 3, nil, 3, nil, 1)), 7)
	assertEq(rob(makeTree(3, 4, 5, 1, 3, nil, 1)), 9)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
