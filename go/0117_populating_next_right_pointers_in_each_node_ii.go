//go:build ignore

// 117. Populating Next Right Pointers in Each Node II (Medium)
// https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
// Run: go run 0117_populating_next_right_pointers_in_each_node_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

type Node struct { // provided by LeetCode
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	for level := root; level != nil; {
		dummy := &Node{} // collects the next level through .Next
		tail := dummy
		for node := level; node != nil; node = node.Next {
			for _, child := range []*Node{node.Left, node.Right} {
				if child != nil {
					tail.Next = child
					tail = child
				}
			}
		}
		level = dummy.Next
	}
	return root
}

func main() {
	n := map[int]*Node{}
	for i := 1; i <= 7; i++ {
		n[i] = &Node{Val: i}
	}
	n[1].Left, n[1].Right = n[2], n[3]
	n[2].Left, n[2].Right, n[3].Right = n[4], n[5], n[7]
	connect(n[1])
	assertEq([]bool{n[2].Next == n[3], n[4].Next == n[5], n[5].Next == n[7], n[7].Next == nil}, []bool{true, true, true, true})
	assertEq(connect(nil) == nil, true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
