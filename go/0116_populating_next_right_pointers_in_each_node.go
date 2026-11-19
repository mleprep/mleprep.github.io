//go:build ignore

// 116. Populating Next Right Pointers in Each Node (Medium)
// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
// Run: go run 0116_populating_next_right_pointers_in_each_node.go lc.go
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
	for leftmost := root; leftmost != nil && leftmost.Left != nil; leftmost = leftmost.Left {
		for node := leftmost; node != nil; node = node.Next { // walk this level with the links built so far
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}

func main() {
	n := make([]*Node, 8) // 1..7 used, heap layout
	for i := range n {
		n[i] = &Node{Val: i}
	}
	for i := 1; i < 4; i++ {
		n[i].Left, n[i].Right = n[2*i], n[2*i+1]
	}
	connect(n[1])
	assertEq([]bool{n[4].Next == n[5], n[5].Next == n[6], n[6].Next == n[7], n[7].Next == nil, n[2].Next == n[3]}, []bool{true, true, true, true, true})
	assertEq(connect(nil) == nil, true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
