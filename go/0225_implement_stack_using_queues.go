//go:build ignore

// 225. Implement Stack using Queues (Easy)
// https://leetcode.com/problems/implement-stack-using-queues/
// Run: go run 0225_implement_stack_using_queues.go lc.go
package main

import (
	"fmt"
	"os"
)

type MyStack struct {
	q []int // front of the queue = top of the stack
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.q = append(this.q, x)
	for i := 0; i < len(this.q)-1; i++ { // rotate the new element to the front
		this.q = append(this.q[1:], this.q[0])
	}
}

func (this *MyStack) Pop() int {
	x := this.q[0]
	this.q = this.q[1:]
	return x
}

func (this *MyStack) Top() int {
	return this.q[0]
}

func (this *MyStack) Empty() bool {
	return len(this.q) == 0
}

func main() {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	assertEq(s.Top(), 2)
	assertEq(s.Pop(), 2)
	assertEq(s.Empty(), false)
	assertEq(s.Pop(), 1)
	assertEq(s.Empty(), true)
	s2 := Constructor()
	s2.Push(5)
	s2.Push(6)
	s2.Push(7)
	assertEq([]int{s2.Pop(), s2.Top(), s2.Pop(), s2.Top()}, []int{7, 6, 6, 5})
	assertEq(s2.Empty(), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
