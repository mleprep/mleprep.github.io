//go:build ignore

// 662. Maximum Width of Binary Tree (Medium)
// https://leetcode.com/problems/maximum-width-of-binary-tree/
// Run: go run 0662_maximum_width_of_binary_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func widthOfBinaryTree(root *TreeNode) int {
	type item struct {
		node *TreeNode
		i    int // heap index within the level
	}
	best, q := 0, []item{{root, 0}}
	for len(q) > 0 {
		base := q[0].i
		best = max(best, q[len(q)-1].i-base+1)
		next := []item{}
		for _, it := range q {
			i := it.i - base // renumber from 0 so indices stay small
			if it.node.Left != nil {
				next = append(next, item{it.node.Left, 2 * i})
			}
			if it.node.Right != nil {
				next = append(next, item{it.node.Right, 2*i + 1})
			}
		}
		q = next
	}
	return best
}

func main() {
	assertEq(widthOfBinaryTree(makeTree(1, 3, 2, 5, 3, nil, 9)), 4)
	assertEq(widthOfBinaryTree(makeTree(1, 3, 2, 5, nil, nil, 9, 6, nil, 7)), 7)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
