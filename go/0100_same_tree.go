//go:build ignore

// 100. Same Tree (Easy)
// https://leetcode.com/problems/same-tree/
// Run: go run 0100_same_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	assertEq(isSameTree(makeTree(1, 2, 3), makeTree(1, 2, 3)), true)
	assertEq(isSameTree(makeTree(1, 2), makeTree(1, nil, 2)), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
