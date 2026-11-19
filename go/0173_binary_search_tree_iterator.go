//go:build ignore

// 173. Binary Search Tree Iterator (Medium)
// https://leetcode.com/problems/binary-search-tree-iterator/
// Run: go run 0173_binary_search_tree_iterator.go lc.go
package main

import (
	"fmt"
	"os"
)

type BSTIterator struct {
	st []*TreeNode // the path of left children still to visit
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{}
	it.pushLeft(root)
	return it
}

func (this *BSTIterator) pushLeft(node *TreeNode) {
	for ; node != nil; node = node.Left {
		this.st = append(this.st, node)
	}
}

func (this *BSTIterator) Next() int {
	node := this.st[len(this.st)-1]
	this.st = this.st[:len(this.st)-1]
	this.pushLeft(node.Right)
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.st) > 0
}

func main() {
	it := Constructor(makeTree(7, 3, 15, nil, nil, 9, 20))
	out := []int{}
	for i := 0; i < 10 && it.HasNext(); i++ {
		out = append(out, it.Next())
	}
	assertEq(out, []int{3, 7, 9, 15, 20})
	it2 := Constructor(makeTree(1))
	assertEq(it2.HasNext(), true)
	assertEq(it2.Next(), 1)
	assertEq(it2.HasNext(), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
