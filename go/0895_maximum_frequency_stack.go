//go:build ignore

// 895. Maximum Frequency Stack (Hard)
// https://leetcode.com/problems/maximum-frequency-stack/
// Run: go run 0895_maximum_frequency_stack.go lc.go
package main

import (
	"fmt"
	"os"
)

type FreqStack struct {
	freq  map[int]int   // value -> current frequency
	group map[int][]int // frequency -> stack of values that reached it
	maxf  int
}

func Constructor() FreqStack {
	return FreqStack{freq: map[int]int{}, group: map[int][]int{}}
}

func (this *FreqStack) Push(val int) {
	f := this.freq[val] + 1
	this.freq[val] = f
	this.group[f] = append(this.group[f], val) // val stays in the lower groups too
	this.maxf = max(this.maxf, f)
}

func (this *FreqStack) Pop() int {
	g := this.group[this.maxf]
	val := g[len(g)-1]
	this.group[this.maxf] = g[:len(g)-1]
	this.freq[val]--
	if len(g) == 1 {
		this.maxf--
	}
	return val
}

func main() {
	fs := Constructor()
	for _, v := range []int{5, 7, 5, 7, 4, 5} {
		fs.Push(v)
	}
	assertEq([]int{fs.Pop(), fs.Pop(), fs.Pop(), fs.Pop()}, []int{5, 7, 5, 4})
	fs.Push(9)
	assertEq([]int{fs.Pop(), fs.Pop(), fs.Pop()}, []int{9, 7, 5})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
