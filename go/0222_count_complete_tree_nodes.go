//go:build ignore

// 222. Count Complete Tree Nodes (Medium)
// https://leetcode.com/problems/count-complete-tree-nodes/
// Run: go run 0222_count_complete_tree_nodes.go lc.go
package main

import (
	"fmt"
	"os"
)

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	hl, hr := 0, 0
	for l := root; l != nil; l = l.Left {
		hl++
	}
	for r := root; r != nil; r = r.Right {
		hr++
	}
	if hl == hr { // perfect tree
		return 1<<hl - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func main() {
	assertEq(countNodes(makeTree(1, 2, 3, 4, 5, 6)), 6)
	vals := []any{}
	for i := 1; i <= 15; i++ {
		vals = append(vals, i)
	}
	assertEq(countNodes(makeTree(vals...)), 15)
	assertEq(countNodes(nil), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
