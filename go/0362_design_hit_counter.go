//go:build ignore

// 362. Design Hit Counter (Medium)
// https://leetcode.com/problems/design-hit-counter/
// Run: go run 0362_design_hit_counter.go lc.go
package main

import (
	"fmt"
	"os"
)

type HitCounter struct {
	q     [][2]int // (timestamp, hits), oldest first
	total int
}

func Constructor() HitCounter {
	return HitCounter{}
}

func (this *HitCounter) Hit(timestamp int) {
	if n := len(this.q); n > 0 && this.q[n-1][0] == timestamp { // same second: one entry
		this.q[n-1][1]++
	} else {
		this.q = append(this.q, [2]int{timestamp, 1})
	}
	this.total++
}

func (this *HitCounter) GetHits(timestamp int) int {
	for len(this.q) > 0 && this.q[0][0] <= timestamp-300 { // older than 5 minutes
		this.total -= this.q[0][1]
		this.q = this.q[1:]
	}
	return this.total
}

func main() {
	hc := Constructor()
	hc.Hit(1)
	hc.Hit(2)
	hc.Hit(3)
	assertEq(hc.GetHits(4), 3)
	hc.Hit(300)
	assertEq([]int{hc.GetHits(300), hc.GetHits(301)}, []int{4, 3})
	hc2 := Constructor()
	for i := 0; i < 1000; i++ {
		hc2.Hit(1)
	}
	assertEq([]int{hc2.GetHits(300), hc2.GetHits(301)}, []int{1000, 0})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
