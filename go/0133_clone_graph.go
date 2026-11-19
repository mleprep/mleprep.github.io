//go:build ignore

// 133. Clone Graph (Medium)
// https://leetcode.com/problems/clone-graph/
// Run: go run 0133_clone_graph.go lc.go
package main

import (
	"fmt"
	"os"
)

type Node struct { // provided by LeetCode
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	clones := map[*Node]*Node{}
	var dfs func(*Node) *Node
	dfs = func(n *Node) *Node {
		if c, ok := clones[n]; ok {
			return c
		}
		c := &Node{Val: n.Val}
		clones[n] = c // register before recursing: the graph has cycles
		for _, nb := range n.Neighbors {
			c.Neighbors = append(c.Neighbors, dfs(nb))
		}
		return c
	}
	if node == nil {
		return nil
	}
	return dfs(node)
}

func main() {
	ns := make([]*Node, 4)
	for i := range ns {
		ns[i] = &Node{Val: i + 1}
	}
	for i, nbs := range [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}} {
		for _, j := range nbs {
			ns[i].Neighbors = append(ns[i].Neighbors, ns[j-1])
		}
	}
	c := cloneGraph(ns[0])
	assertEq(c != ns[0], true)
	assertEq([]int{c.Neighbors[0].Val, c.Neighbors[1].Val}, []int{2, 4})
	assertEq(c.Neighbors[0].Neighbors[0] == c, true) // cycle preserved
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
