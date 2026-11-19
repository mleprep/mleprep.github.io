//go:build ignore

// 138. Copy List with Random Pointer (Medium)
// https://leetcode.com/problems/copy-list-with-random-pointer/
// Run: go run 0138_copy_list_with_random_pointer.go lc.go
package main

import (
	"fmt"
	"os"
)

type Node struct { // provided by LeetCode
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	clone := map[*Node]*Node{nil: nil}
	for cur := head; cur != nil; cur = cur.Next {
		clone[cur] = &Node{Val: cur.Val}
	}
	for cur := head; cur != nil; cur = cur.Next {
		clone[cur].Next = clone[cur.Next]
		clone[cur].Random = clone[cur.Random]
	}
	return clone[head]
}

func main() {
	spec := [][2]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}} // (val, random index or -1)
	nodes := make([]*Node, len(spec))
	for i, s := range spec {
		nodes[i] = &Node{Val: s[0]}
	}
	for i, s := range spec {
		if i+1 < len(nodes) {
			nodes[i].Next = nodes[i+1]
		}
		if s[1] >= 0 {
			nodes[i].Random = nodes[s[1]]
		}
	}
	describe := func(head *Node) [][2]int { // the copy's spec; nil if it reuses an original node
		orig := map[*Node]bool{}
		for _, n := range nodes {
			orig[n] = true
		}
		var list []*Node
		for c := head; c != nil; c = c.Next {
			if orig[c] {
				return nil
			}
			list = append(list, c)
		}
		pos := map[*Node]int{}
		for i, c := range list {
			pos[c] = i
		}
		out := [][2]int{}
		for _, c := range list {
			r := -1
			if c.Random != nil {
				r = pos[c.Random]
			}
			out = append(out, [2]int{c.Val, r})
		}
		return out
	}
	assertEq(describe(copyRandomList(nodes[0])), spec)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
