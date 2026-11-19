//go:build ignore

// 113. Path Sum II (Medium)
// https://leetcode.com/problems/path-sum-ii/
// Run: go run 0113_path_sum_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func pathSum(root *TreeNode, targetSum int) [][]int {
	res, path := [][]int{}, []int{}
	var dfs func(node *TreeNode, rest int)
	dfs = func(node *TreeNode, rest int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		rest -= node.Val
		if node.Left == nil && node.Right == nil && rest == 0 {
			res = append(res, append([]int{}, path...)) // copy: path keeps changing
		}
		dfs(node.Left, rest)
		dfs(node.Right, rest)
		path = path[:len(path)-1]
	}
	dfs(root, targetSum)
	return res
}

func main() {
	t := makeTree(5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1)
	assertUnordered(pathSum(t, 22), [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}})
	assertEq(pathSum(makeTree(1, 2, 3), 5), [][]int{})
	assertEq(pathSum(makeTree(1, 2), 0), [][]int{})
	assertEq(pathSum(makeTree(-2, nil, -3), -5), [][]int{{-2, -3}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
