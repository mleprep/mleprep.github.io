//go:build ignore

// 706. Design HashMap (Easy)
// https://leetcode.com/problems/design-hashmap/
// Run: go run 0706_design_hashmap.go lc.go
package main

import (
	"fmt"
	"os"
)

type MyHashMap struct {
	b [][][2]int // separate chaining: buckets of (key, value)
}

const buckets = 1009 // a prime number of buckets

func Constructor() MyHashMap {
	return MyHashMap{b: make([][][2]int, buckets)}
}

func (this *MyHashMap) Put(key int, value int) {
	b := this.b[key%buckets]
	for i := range b {
		if b[i][0] == key {
			b[i][1] = value
			return
		}
	}
	this.b[key%buckets] = append(b, [2]int{key, value})
}

func (this *MyHashMap) Get(key int) int {
	for _, kv := range this.b[key%buckets] {
		if kv[0] == key {
			return kv[1]
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	b := this.b[key%buckets]
	for i := range b {
		if b[i][0] == key {
			this.b[key%buckets] = append(b[:i], b[i+1:]...)
			return
		}
	}
}

func main() {
	hm := Constructor()
	hm.Put(1, 1)
	hm.Put(2, 2)
	assertEq([]int{hm.Get(1), hm.Get(3)}, []int{1, -1})
	hm.Put(2, 1)
	assertEq(hm.Get(2), 1)
	hm.Remove(2)
	assertEq(hm.Get(2), -1)
	hm.Put(1009+1, 7) // same bucket as 1
	assertEq([]int{hm.Get(1), hm.Get(1010)}, []int{1, 7})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
