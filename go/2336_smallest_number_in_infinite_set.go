//go:build ignore

// 2336. Smallest Number in Infinite Set (Medium)
// https://leetcode.com/problems/smallest-number-in-infinite-set/
// Run: go run 2336_smallest_number_in_infinite_set.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"os"
)

type IntHeap []int // min-heap

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type SmallestInfiniteSet struct {
	next   int          // every number >= next is in the set
	back   IntHeap      // numbers < next that were added back
	inBack map[int]bool // ...and their membership
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{next: 1, inBack: map[int]bool{}}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.back.Len() > 0 {
		x := heap.Pop(&this.back).(int)
		delete(this.inBack, x)
		return x
	}
	this.next++
	return this.next - 1
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num < this.next && !this.inBack[num] {
		this.inBack[num] = true
		heap.Push(&this.back, num)
	}
}

func main() {
	s := Constructor()
	s.AddBack(2)
	assertEq([]int{s.PopSmallest(), s.PopSmallest(), s.PopSmallest()}, []int{1, 2, 3})
	s.AddBack(1)
	assertEq([]int{s.PopSmallest(), s.PopSmallest(), s.PopSmallest()}, []int{1, 4, 5})
	s2 := Constructor()
	s2.PopSmallest()
	s2.PopSmallest()
	s2.AddBack(1)
	s2.AddBack(1)
	assertEq([]int{s2.PopSmallest(), s2.PopSmallest()}, []int{1, 3})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
