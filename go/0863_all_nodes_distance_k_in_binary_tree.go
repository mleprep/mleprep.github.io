//go:build ignore

// 863. All Nodes Distance K in Binary Tree (Medium)
// https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/
// Run: go run 0863_all_nodes_distance_k_in_binary_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parent := map[*TreeNode]*TreeNode{}
	st := []*TreeNode{root}
	for len(st) > 0 { // record every node's parent, then BFS outward from target
		node := st[len(st)-1]
		st = st[:len(st)-1]
		for _, c := range []*TreeNode{node.Left, node.Right} {
			if c != nil {
				parent[c] = node
				st = append(st, c)
			}
		}
	}
	q, seen := []*TreeNode{target}, map[*TreeNode]bool{target: true}
	for ; k > 0 && len(q) > 0; k-- {
		next := []*TreeNode{}
		for _, node := range q {
			for _, nb := range []*TreeNode{node.Left, node.Right, parent[node]} {
				if nb != nil && !seen[nb] {
					seen[nb] = true
					next = append(next, nb)
				}
			}
		}
		q = next
	}
	res := []int{}
	for _, node := range q {
		res = append(res, node.Val)
	}
	return res
}

func main() {
	t := makeTree(3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4)
	assertUnordered(distanceK(t, findNode(t, 5), 2), []int{1, 4, 7})
	assertEq(distanceK(t, findNode(t, 5), 0), []int{5})
	assertUnordered(distanceK(t, findNode(t, 7), 3), []int{3, 6})
	t1 := makeTree(1)
	assertEq(distanceK(t1, t1, 3), []int{})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
