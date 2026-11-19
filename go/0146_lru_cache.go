//go:build ignore

// 146. LRU Cache (Medium)
// https://leetcode.com/problems/lru-cache/
// Run: go run 0146_lru_cache.go lc.go
package main

import (
	"fmt"
	"os"
)

type lruNode struct {
	key, val   int
	prev, next *lruNode
}

type LRUCache struct {
	cap        int
	nodes      map[int]*lruNode
	head, tail *lruNode // sentinels: head.next is the most recent, tail.prev the least
}

func Constructor(capacity int) LRUCache {
	head, tail := &lruNode{}, &lruNode{}
	head.next, tail.prev = tail, head
	return LRUCache{cap: capacity, nodes: map[int]*lruNode{}, head: head, tail: tail}
}

func (this *LRUCache) unlink(n *lruNode) {
	n.prev.next, n.next.prev = n.next, n.prev
}

func (this *LRUCache) pushFront(n *lruNode) {
	n.prev, n.next = this.head, this.head.next
	this.head.next.prev = n
	this.head.next = n
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.nodes[key]
	if !ok {
		return -1
	}
	this.unlink(n) // used: move to the front
	this.pushFront(n)
	return n.val
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.nodes[key]; ok {
		n.val = value
		this.unlink(n)
		this.pushFront(n)
		return
	}
	n := &lruNode{key: key, val: value}
	this.nodes[key] = n
	this.pushFront(n)
	if len(this.nodes) > this.cap { // evict the least recently used
		lru := this.tail.prev
		this.unlink(lru)
		delete(this.nodes, lru.key)
	}
}

func main() {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	assertEq(c.Get(1), 1)
	c.Put(3, 3) // evicts 2
	assertEq(c.Get(2), -1)
	c.Put(4, 4) // evicts 1
	assertEq([]int{c.Get(1), c.Get(3), c.Get(4)}, []int{-1, 3, 4})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
