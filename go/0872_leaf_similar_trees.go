//go:build ignore

// 872. Leaf-Similar Trees (Easy)
// https://leetcode.com/problems/leaf-similar-trees/
// Run: go run 0872_leaf_similar_trees.go lc.go
package main

import (
	"fmt"
	"os"
)

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leaves func(node *TreeNode, out []int) []int
	leaves = func(node *TreeNode, out []int) []int {
		if node == nil {
			return out
		}
		if node.Left == nil && node.Right == nil {
			return append(out, node.Val)
		}
		return leaves(node.Right, leaves(node.Left, out))
	}
	a, b := leaves(root1, nil), leaves(root2, nil)
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	a := makeTree(3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4)
	b := makeTree(3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8)
	assertEq(leafSimilar(a, b), true)
	assertEq(leafSimilar(makeTree(1, 2, 3), makeTree(1, 3, 2)), false)
	assertEq(leafSimilar(makeTree(1), makeTree(2, 1)), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
