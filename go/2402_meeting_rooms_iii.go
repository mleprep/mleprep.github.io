//go:build ignore

// 2402. Meeting Rooms III (Hard)
// https://leetcode.com/problems/meeting-rooms-iii/
// Run: go run 2402_meeting_rooms_iii.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type IntHeap []int // free rooms, lowest number first

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

type busyHeap [][2]int // (free-at time, room): earliest first, then the lowest room

func (h busyHeap) Len() int { return len(h) }
func (h busyHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0] || h[i][0] == h[j][0] && h[i][1] < h[j][1]
}
func (h busyHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *busyHeap) Push(x any)   { *h = append(*h, x.([2]int)) }
func (h *busyHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][0] < meetings[j][0] })
	free, busy, count := &IntHeap{}, &busyHeap{}, make([]int, n)
	for r := 0; r < n; r++ {
		heap.Push(free, r)
	}
	for _, m := range meetings {
		s, e := m[0], m[1]
		for busy.Len() > 0 && (*busy)[0][0] <= s { // release rooms free by s
			heap.Push(free, heap.Pop(busy).([2]int)[1])
		}
		var room int
		if free.Len() > 0 {
			room = heap.Pop(free).(int)
			heap.Push(busy, [2]int{e, room})
		} else { // wait for the earliest room; the meeting keeps its duration
			top := heap.Pop(busy).([2]int)
			room = top[1]
			heap.Push(busy, [2]int{top[0] + e - s, room})
		}
		count[room]++
	}
	best := 0
	for r := range count {
		if count[r] > count[best] {
			best = r
		}
	}
	return best
}

func main() {
	assertEq(mostBooked(2, [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}), 0)
	assertEq(mostBooked(3, [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}), 1)
	assertEq(mostBooked(1, [][]int{{0, 1}}), 0)
	assertEq(mostBooked(4, [][]int{{18, 19}, {3, 12}, {17, 19}, {2, 13}, {7, 10}}), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
