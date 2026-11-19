//go:build ignore

// 637. Average of Levels in Binary Tree (Easy)
// https://leetcode.com/problems/average-of-levels-in-binary-tree/
// Run: go run 0637_average_of_levels_in_binary_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func averageOfLevels(root *TreeNode) []float64 {
	res, q := []float64{}, []*TreeNode{root}
	for len(q) > 0 {
		sum, next := 0, []*TreeNode{}
		for _, n := range q {
			sum += n.Val
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		res = append(res, float64(sum)/float64(len(q)))
		q = next
	}
	return res
}

func main() {
	assertEq(averageOfLevels(makeTree(3, 9, 20, nil, nil, 15, 7)), []float64{3, 14.5, 11})
	assertEq(averageOfLevels(makeTree(1)), []float64{1})
	assertEq(averageOfLevels(makeTree(2147483647, 2147483647, 2147483647)), []float64{2147483647, 2147483647})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
