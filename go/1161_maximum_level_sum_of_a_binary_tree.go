//go:build ignore

// 1161. Maximum Level Sum of a Binary Tree (Medium)
// https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
// Run: go run 1161_maximum_level_sum_of_a_binary_tree.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func maxLevelSum(root *TreeNode) int {
	best, bestLevel := math.MinInt, 0
	q := []*TreeNode{root}
	for level := 1; len(q) > 0; level++ {
		total, next := 0, []*TreeNode{}
		for _, n := range q {
			total += n.Val
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		if total > best { // strictly greater: the smallest level wins ties
			best, bestLevel = total, level
		}
		q = next
	}
	return bestLevel
}

func main() {
	assertEq(maxLevelSum(makeTree(1, 7, 0, 7, -8, nil, nil)), 2)
	assertEq(maxLevelSum(makeTree(989, nil, 10250, 98693, -89388, nil, nil, nil, -32127)), 2)
	assertEq(maxLevelSum(makeTree(-1, -2, -3)), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
