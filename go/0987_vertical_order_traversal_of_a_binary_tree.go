//go:build ignore

// 987. Vertical Order Traversal of a Binary Tree (Hard)
// https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/
// Run: go run 0987_vertical_order_traversal_of_a_binary_tree.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func verticalTraversal(root *TreeNode) [][]int {
	cells := [][3]int{} // (col, row, val)
	var dfs func(node *TreeNode, r, c int)
	dfs = func(node *TreeNode, r, c int) {
		if node != nil {
			cells = append(cells, [3]int{c, r, node.Val})
			dfs(node.Left, r+1, c-1)
			dfs(node.Right, r+1, c+1)
		}
	}
	dfs(root, 0, 0)
	sort.Slice(cells, func(i, j int) bool {
		a, b := cells[i], cells[j]
		if a[0] != b[0] {
			return a[0] < b[0]
		}
		if a[1] != b[1] {
			return a[1] < b[1]
		}
		return a[2] < b[2]
	})
	res := [][]int{}
	for i, cell := range cells {
		if i == 0 || cell[0] != cells[i-1][0] {
			res = append(res, []int{})
		}
		res[len(res)-1] = append(res[len(res)-1], cell[2])
	}
	return res
}

func main() {
	assertEq(verticalTraversal(makeTree(3, 9, 20, nil, nil, 15, 7)), [][]int{{9}, {3, 15}, {20}, {7}})
	assertEq(verticalTraversal(makeTree(1, 2, 3, 4, 6, 5, 7)), [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
