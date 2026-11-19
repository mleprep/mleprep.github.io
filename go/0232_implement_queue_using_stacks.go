//go:build ignore

// 232. Implement Queue using Stacks (Easy)
// https://leetcode.com/problems/implement-queue-using-stacks/
// Run: go run 0232_implement_queue_using_stacks.go lc.go
package main

import (
	"fmt"
	"os"
)

type MyQueue struct {
	in, out []int // push onto in; pop from out, refilled (reversed) only when empty
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) move() {
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			this.out = append(this.out, this.in[len(this.in)-1])
			this.in = this.in[:len(this.in)-1]
		}
	}
}

func (this *MyQueue) Pop() int {
	this.move()
	x := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return x
}

func (this *MyQueue) Peek() int {
	this.move()
	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

func main() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	assertEq([]any{q.Peek(), q.Pop(), q.Empty()}, []any{1, 1, false})
	q.Push(3)
	assertEq([]any{q.Pop(), q.Peek(), q.Pop(), q.Empty()}, []any{2, 3, 3, true})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
