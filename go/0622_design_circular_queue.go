//go:build ignore

// 622. Design Circular Queue (Medium)
// https://leetcode.com/problems/design-circular-queue/
// Run: go run 0622_design_circular_queue.go lc.go
package main

import (
	"fmt"
	"os"
)

type MyCircularQueue struct {
	a       []int
	head, n int // index of the front, number of items
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{a: make([]int, k)}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.a[(this.head+this.n)%len(this.a)] = value
	this.n++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % len(this.a)
	this.n--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.a[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.a[(this.head+this.n-1)%len(this.a)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.n == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.n == len(this.a)
}

func main() {
	cq := Constructor(3)
	assertEq([]bool{cq.EnQueue(1), cq.EnQueue(2), cq.EnQueue(3), cq.EnQueue(4)}, []bool{true, true, true, false})
	assertEq([]any{cq.Rear(), cq.IsFull(), cq.DeQueue(), cq.EnQueue(4), cq.Rear(), cq.Front()}, []any{3, true, true, true, 4, 2})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
