//go:build ignore

// 460. LFU Cache (Hard)
// https://leetcode.com/problems/lfu-cache/
// Run: go run 0460_lfu_cache.go lc.go
package main

import (
	"container/list"
	"fmt"
	"os"
)

type lfuEntry struct {
	key, val, freq int
}

type LFUCache struct {
	cap, minFreq int
	items        map[int]*list.Element // key -> its element in lists[freq]
	lists        map[int]*list.List    // freq -> entries, most recently used at the front
}

func Constructor(capacity int) LFUCache {
	return LFUCache{cap: capacity, items: map[int]*list.Element{}, lists: map[int]*list.List{}}
}

func (this *LFUCache) touch(el *list.Element) { // one more use: move to the next frequency list
	e := el.Value.(*lfuEntry)
	old := this.lists[e.freq]
	old.Remove(el)
	if old.Len() == 0 {
		delete(this.lists, e.freq)
		if this.minFreq == e.freq {
			this.minFreq++
		}
	}
	e.freq++
	if this.lists[e.freq] == nil {
		this.lists[e.freq] = list.New()
	}
	this.items[e.key] = this.lists[e.freq].PushFront(e)
}

func (this *LFUCache) Get(key int) int {
	el, ok := this.items[key]
	if !ok {
		return -1
	}
	e := el.Value.(*lfuEntry)
	this.touch(el)
	return e.val
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	if el, ok := this.items[key]; ok {
		e := el.Value.(*lfuEntry)
		e.val = value
		this.touch(el)
		return
	}
	if len(this.items) == this.cap { // evict the LRU entry among the least frequent
		l := this.lists[this.minFreq]
		victim := l.Back()
		l.Remove(victim)
		if l.Len() == 0 {
			delete(this.lists, this.minFreq)
		}
		delete(this.items, victim.Value.(*lfuEntry).key)
	}
	if this.lists[1] == nil {
		this.lists[1] = list.New()
	}
	this.items[key] = this.lists[1].PushFront(&lfuEntry{key: key, val: value, freq: 1})
	this.minFreq = 1
}

func main() {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	assertEq(c.Get(1), 1)
	c.Put(3, 3) // evicts 2 (freq 1)
	assertEq([]int{c.Get(2), c.Get(3)}, []int{-1, 3})
	c.Put(4, 4) // 1 and 3 both have freq 2: evict the least recent of them, 1
	assertEq([]int{c.Get(1), c.Get(3), c.Get(4)}, []int{-1, 3, 4})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
