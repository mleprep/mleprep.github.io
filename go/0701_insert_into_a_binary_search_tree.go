//go:build ignore

// 701. Insert into a Binary Search Tree (Medium)
// https://leetcode.com/problems/insert-into-a-binary-search-tree/
// Run: go run 0701_insert_into_a_binary_search_tree.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	if root == nil {
		return node
	}
	for cur := root; ; {
		if val < cur.Val {
			if cur.Left == nil {
				cur.Left = node
				return root
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = node
				return root
			}
			cur = cur.Right
		}
	}
}

func main() {
	var inorder func(n *TreeNode) []int
	inorder = func(n *TreeNode) []int {
		if n == nil {
			return []int{}
		}
		return append(append(inorder(n.Left), n.Val), inorder(n.Right)...)
	}
	for _, tc := range []struct {
		vals []any
		v    int
	}{{[]any{4, 2, 7, 1, 3}, 5}, {[]any{40, 20, 60, 10, 30, 50, 70}, 25}, {[]any{}, 5}, {[]any{4, 2, 7, 1, 3, nil, nil, nil, nil, nil, nil}, 0}} {
		want := []int{tc.v}
		for _, x := range tc.vals {
			if x != nil {
				want = append(want, x.(int))
			}
		}
		sort.Ints(want)
		assertEq(inorder(insertIntoBST(makeTree(tc.vals...), tc.v)), want) // still a BST holding every value
	}
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
