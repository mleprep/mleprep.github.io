//go:build ignore

// 437. Path Sum III (Medium)
// https://leetcode.com/problems/path-sum-iii/
// Run: go run 0437_path_sum_iii.go lc.go
package main

import (
	"fmt"
	"os"
)

func pathSum(root *TreeNode, targetSum int) int {
	seen := map[int]int{0: 1} // prefix sums on the current root-to-node path
	var dfs func(node *TreeNode, cur int) int
	dfs = func(node *TreeNode, cur int) int {
		if node == nil {
			return 0
		}
		cur += node.Val
		cnt := seen[cur-targetSum]
		seen[cur]++
		cnt += dfs(node.Left, cur) + dfs(node.Right, cur)
		seen[cur]-- // backtrack
		return cnt
	}
	return dfs(root, 0)
}

func main() {
	assertEq(pathSum(makeTree(10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1), 8), 3)
	assertEq(pathSum(makeTree(5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1), 22), 3)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
