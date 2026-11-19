//go:build ignore

// 341. Flatten Nested List Iterator (Medium)
// https://leetcode.com/problems/flatten-nested-list-iterator/
// Run: go run 0341_flatten_nested_list_iterator.go lc.go
package main

import (
	"fmt"
	"os"
)

type NestedInteger struct { // provided by LeetCode (a minimal stand-in)
	isInt bool
	val   int
	list  []*NestedInteger
}

func (this NestedInteger) IsInteger() bool { return this.isInt }

func (this NestedInteger) GetInteger() int { return this.val }

func (this NestedInteger) GetList() []*NestedInteger { return this.list }

type NestedIterator struct {
	st []*NestedInteger // reversed: the next element is on top
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	st := make([]*NestedInteger, len(nestedList))
	for i, x := range nestedList {
		st[len(nestedList)-1-i] = x
	}
	return &NestedIterator{st: st}
}

func (this *NestedIterator) Next() int {
	top := this.st[len(this.st)-1] // HasNext() left an integer on top
	this.st = this.st[:len(this.st)-1]
	return top.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	for len(this.st) > 0 && !this.st[len(this.st)-1].IsInteger() { // unpack lists lazily
		top := this.st[len(this.st)-1]
		this.st = this.st[:len(this.st)-1]
		l := top.GetList()
		for i := len(l) - 1; i >= 0; i-- {
			this.st = append(this.st, l[i])
		}
	}
	return len(this.st) > 0
}

func main() {
	var wrap func(x any) *NestedInteger
	wrap = func(x any) *NestedInteger {
		if v, ok := x.(int); ok {
			return &NestedInteger{isInt: true, val: v}
		}
		n := &NestedInteger{}
		for _, y := range x.([]any) {
			n.list = append(n.list, wrap(y))
		}
		return n
	}
	flatten := func(items []any) []int {
		nested := []*NestedInteger{}
		for _, x := range items {
			nested = append(nested, wrap(x))
		}
		it, out := Constructor(nested), []int{}
		for it.HasNext() {
			out = append(out, it.Next())
		}
		return out
	}
	assertEq(flatten([]any{[]any{1, 1}, 2, []any{1, 1}}), []int{1, 1, 2, 1, 1})
	assertEq(flatten([]any{1, []any{4, []any{6}}, []any{[]any{}}}), []int{1, 4, 6})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
