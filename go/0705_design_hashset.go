//go:build ignore

// 705. Design HashSet (Easy)
// https://leetcode.com/problems/design-hashset/
// Run: go run 0705_design_hashset.go lc.go
package main

import (
	"fmt"
	"os"
)

type MyHashSet struct {
	buckets [][]int // separate chaining
}

const size = 1009 // a prime number of buckets

func Constructor() MyHashSet {
	return MyHashSet{buckets: make([][]int, size)}
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		this.buckets[key%size] = append(this.buckets[key%size], key)
	}
}

func (this *MyHashSet) Remove(key int) {
	b := this.buckets[key%size]
	for i, k := range b {
		if k == key {
			this.buckets[key%size] = append(b[:i], b[i+1:]...)
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	for _, k := range this.buckets[key%size] {
		if k == key {
			return true
		}
	}
	return false
}

func main() {
	hs := Constructor()
	hs.Add(1)
	hs.Add(2)
	assertEq([]bool{hs.Contains(1), hs.Contains(3)}, []bool{true, false})
	hs.Add(2)
	assertEq(hs.Contains(2), true)
	hs.Remove(2)
	assertEq(hs.Contains(2), false)
	hs.Add(0) // 0, 1009 and 2018 share a bucket
	hs.Add(1009)
	hs.Add(2018)
	hs.Remove(1009)
	assertEq([]bool{hs.Contains(0), hs.Contains(1009), hs.Contains(2018)}, []bool{true, false, true})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
