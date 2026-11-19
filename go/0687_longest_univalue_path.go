//go:build ignore

// 687. Longest Univalue Path (Medium)
// https://leetcode.com/problems/longest-univalue-path/
// Run: go run 0687_longest_univalue_path.go lc.go
package main

import (
	"fmt"
	"os"
)

func longestUnivaluePath(root *TreeNode) int {
	best := 0
	var arm func(*TreeNode) int // longest same-value path going down from node
	arm = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := arm(node.Left), arm(node.Right)
		if node.Left != nil && node.Left.Val == node.Val {
			l++
		} else {
			l = 0
		}
		if node.Right != nil && node.Right.Val == node.Val {
			r++
		} else {
			r = 0
		}
		best = max(best, l+r)
		return max(l, r)
	}
	arm(root)
	return best
}

func main() {
	assertEq(longestUnivaluePath(makeTree(5, 4, 5, 1, 1, nil, 5)), 2)
	assertEq(longestUnivaluePath(makeTree(1, 4, 5, 4, 4, nil, 5)), 2)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
