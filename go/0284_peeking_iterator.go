//go:build ignore

// 284. Peeking Iterator (Medium)
// https://leetcode.com/problems/peeking-iterator/
// Run: go run 0284_peeking_iterator.go lc.go
package main

import (
	"fmt"
	"os"
)

type Iterator struct { // provided by LeetCode (a minimal stand-in)
	nums []int
	i    int
}

func (this *Iterator) hasNext() bool {
	return this.i < len(this.nums)
}

func (this *Iterator) next() int {
	this.i++
	return this.nums[this.i-1]
}

type PeekingIterator struct {
	iter    *Iterator
	has     bool
	nextVal int // cached next element
}

func Constructor(iter *Iterator) *PeekingIterator {
	p := &PeekingIterator{iter: iter}
	p.advance()
	return p
}

func (this *PeekingIterator) advance() {
	this.has = this.iter.hasNext()
	if this.has {
		this.nextVal = this.iter.next()
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.has
}

func (this *PeekingIterator) next() int {
	v := this.nextVal
	this.advance()
	return v
}

func (this *PeekingIterator) peek() int {
	return this.nextVal
}

func main() {
	pi := Constructor(&Iterator{nums: []int{1, 2, 3}})
	assertEq([]any{pi.next(), pi.peek(), pi.next(), pi.next(), pi.hasNext()}, []any{1, 2, 2, 3, false})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
