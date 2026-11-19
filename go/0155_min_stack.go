//go:build ignore

// 155. Min Stack (Medium)
// https://leetcode.com/problems/min-stack/
// Run: go run 0155_min_stack.go lc.go
package main

import (
	"fmt"
	"os"
)

type MinStack struct {
	st [][2]int // (val, min so far)
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	m := val
	if n := len(this.st); n > 0 {
		m = min(m, this.st[n-1][1])
	}
	this.st = append(this.st, [2]int{val, m})
}

func (this *MinStack) Pop() {
	this.st = this.st[:len(this.st)-1]
}

func (this *MinStack) Top() int {
	return this.st[len(this.st)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.st[len(this.st)-1][1]
}

func main() {
	m := Constructor()
	m.Push(-2)
	m.Push(0)
	m.Push(-3)
	assertEq(m.GetMin(), -3)
	m.Pop()
	assertEq(m.Top(), 0)
	assertEq(m.GetMin(), -2)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
