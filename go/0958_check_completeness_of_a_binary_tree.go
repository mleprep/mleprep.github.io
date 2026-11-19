//go:build ignore

// 958. Check Completeness of a Binary Tree (Medium)
// https://leetcode.com/problems/check-completeness-of-a-binary-tree/
// Run: go run 0958_check_completeness_of_a_binary_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func isCompleteTree(root *TreeNode) bool {
	q, seenNil := []*TreeNode{root}, false
	for len(q) > 0 { // BFS including nils: no node may come after a nil
		node := q[0]
		q = q[1:]
		if node == nil {
			seenNil = true
			continue
		}
		if seenNil {
			return false
		}
		q = append(q, node.Left, node.Right)
	}
	return true
}

func main() {
	assertEq(isCompleteTree(makeTree(1, 2, 3, 4, 5, 6)), true)
	assertEq(isCompleteTree(makeTree(1, 2, 3, 4, 5, nil, 7)), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
