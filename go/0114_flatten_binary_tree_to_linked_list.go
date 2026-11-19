//go:build ignore

// 114. Flatten Binary Tree to Linked List (Medium)
// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
// Run: go run 0114_flatten_binary_tree_to_linked_list.go lc.go
package main

import (
	"fmt"
	"os"
)

func flatten(root *TreeNode) {
	for cur := root; cur != nil; cur = cur.Right {
		if cur.Left != nil {
			p := cur.Left // rightmost node of the left subtree
			for p.Right != nil {
				p = p.Right
			}
			p.Right = cur.Right
			cur.Right, cur.Left = cur.Left, nil
		}
	}
}

func main() {
	t := makeTree(1, 2, 5, 3, 4, nil, 6)
	flatten(t)
	assertEq(t, []any{1, nil, 2, nil, 3, nil, 4, nil, 5, nil, 6})
	t2 := makeTree(0)
	flatten(t2)
	assertEq(t2, []any{0})
	flatten(nil)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
