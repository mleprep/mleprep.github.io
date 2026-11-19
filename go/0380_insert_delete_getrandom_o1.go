//go:build ignore

// 380. Insert Delete GetRandom O(1) (Medium)
// https://leetcode.com/problems/insert-delete-getrandom-o1/
// Run: go run 0380_insert_delete_getrandom_o1.go lc.go
package main

import (
	"fmt"
	"math/rand"
	"os"
)

type RandomizedSet struct {
	vals []int
	idx  map[int]int // value -> position in vals
}

func Constructor() RandomizedSet {
	return RandomizedSet{idx: map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.idx[val]; ok {
		return false
	}
	this.idx[val] = len(this.vals)
	this.vals = append(this.vals, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	i, ok := this.idx[val]
	if !ok {
		return false
	}
	last := this.vals[len(this.vals)-1] // move the last value into the hole: O(1)
	this.vals[i] = last
	this.idx[last] = i
	this.vals = this.vals[:len(this.vals)-1]
	delete(this.idx, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.vals[rand.Intn(len(this.vals))]
}

func main() {
	rs := Constructor()
	assertEq([]bool{rs.Insert(1), rs.Remove(2), rs.Insert(2)}, []bool{true, false, true})
	r := rs.GetRandom()
	assertEq(r == 1 || r == 2, true)
	assertEq([]bool{rs.Remove(1), rs.Insert(2)}, []bool{true, false})
	assertEq(rs.GetRandom(), 2)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
