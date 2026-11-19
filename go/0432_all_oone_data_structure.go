//go:build ignore

// 432. All O`one Data Structure (Hard)
// https://leetcode.com/problems/all-oone-data-structure/
// Run: go run 0432_all_oone_data_structure.go lc.go
package main

import (
	"fmt"
	"os"
)

type bucket struct {
	cnt        int
	keys       map[string]bool
	prev, next *bucket
}

type AllOne struct {
	head, tail *bucket            // sentinels; counts increase from head to tail
	where      map[string]*bucket // key -> the bucket holding it
}

func Constructor() AllOne {
	head, tail := &bucket{}, &bucket{}
	head.next, tail.prev = tail, head
	return AllOne{head: head, tail: tail, where: map[string]*bucket{}}
}

func (this *AllOne) after(node *bucket, cnt int) *bucket {
	b := &bucket{cnt: cnt, keys: map[string]bool{}, prev: node, next: node.next}
	node.next.prev = b
	node.next = b
	return b
}

func (this *AllOne) dropIfEmpty(b *bucket) {
	if len(b.keys) == 0 {
		b.prev.next, b.next.prev = b.next, b.prev
	}
}

func (this *AllOne) Inc(key string) {
	cur, ok := this.where[key]
	if !ok {
		cur = this.head // count 0
	}
	nxt := cur.next
	if nxt == this.tail || nxt.cnt != cur.cnt+1 {
		nxt = this.after(cur, cur.cnt+1)
	}
	nxt.keys[key] = true
	this.where[key] = nxt
	if cur != this.head {
		delete(cur.keys, key)
		this.dropIfEmpty(cur)
	}
}

func (this *AllOne) Dec(key string) {
	cur := this.where[key]
	if cur.cnt == 1 {
		delete(this.where, key)
	} else {
		prv := cur.prev
		if prv == this.head || prv.cnt != cur.cnt-1 {
			prv = this.after(cur.prev, cur.cnt-1)
		}
		prv.keys[key] = true
		this.where[key] = prv
	}
	delete(cur.keys, key)
	this.dropIfEmpty(cur)
}

func anyKey(b *bucket) string {
	for k := range b.keys {
		return k
	}
	return ""
}

func (this *AllOne) GetMaxKey() string {
	if this.tail.prev == this.head {
		return ""
	}
	return anyKey(this.tail.prev)
}

func (this *AllOne) GetMinKey() string {
	if this.head.next == this.tail {
		return ""
	}
	return anyKey(this.head.next)
}

func main() {
	a := Constructor()
	a.Inc("hello")
	a.Inc("hello")
	assertEq([]string{a.GetMaxKey(), a.GetMinKey()}, []string{"hello", "hello"})
	a.Inc("leet")
	assertEq([]string{a.GetMaxKey(), a.GetMinKey()}, []string{"hello", "leet"})
	a.Dec("hello")
	a.Dec("hello")
	assertEq([]string{a.GetMaxKey(), a.GetMinKey()}, []string{"leet", "leet"})
	a.Dec("leet")
	assertEq([]string{a.GetMaxKey(), a.GetMinKey()}, []string{"", ""})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
