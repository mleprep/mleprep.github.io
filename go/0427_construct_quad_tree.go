//go:build ignore

// 427. Construct Quad Tree (Medium)
// https://leetcode.com/problems/construct-quad-tree/
// Run: go run 0427_construct_quad_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

type Node struct { // provided by LeetCode
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	var build func(r, c, n int) *Node
	build = func(r, c, n int) *Node {
		if n == 1 {
			return &Node{Val: grid[r][c] == 1, IsLeaf: true}
		}
		h := n / 2
		kids := []*Node{build(r, c, h), build(r, c+h, h), build(r+h, c, h), build(r+h, c+h, h)}
		merge := true // four leaves with the same value become one leaf
		for _, k := range kids {
			merge = merge && k.IsLeaf && k.Val == kids[0].Val
		}
		if merge {
			return &Node{Val: kids[0].Val, IsLeaf: true}
		}
		return &Node{Val: true, TopLeft: kids[0], TopRight: kids[1], BottomLeft: kids[2], BottomRight: kids[3]}
	}
	return build(0, 0, len(grid))
}

func main() {
	quad := func(node *Node) string { // level order: "*" = internal node, "0"/"1" = leaf value
		out, q := "", []*Node{node}
		for len(q) > 0 {
			n := q[0]
			q = q[1:]
			if n == nil {
				return out + "?"
			}
			switch {
			case !n.IsLeaf:
				out += "*"
				q = append(q, n.TopLeft, n.TopRight, n.BottomLeft, n.BottomRight)
			case n.Val:
				out += "1"
			default:
				out += "0"
			}
		}
		return out
	}
	assertEq(quad(construct([][]int{{0, 1}, {1, 0}})), "*0110")
	assertEq(quad(construct([][]int{{1}})), "1")
	assertEq(quad(construct([][]int{{1, 1}, {1, 1}})), "1")
	assertEq(quad(construct([][]int{{1, 1, 0, 0}, {1, 1, 0, 0}, {1, 1, 1, 1}, {1, 1, 1, 1}})), "*1011")
	assertEq(quad(construct([][]int{{1, 1, 0, 0}, {1, 0, 0, 0}, {1, 1, 1, 1}, {1, 1, 1, 1}})), "**0111110")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
