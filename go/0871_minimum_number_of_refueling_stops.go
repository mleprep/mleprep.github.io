//go:build ignore

// 871. Minimum Number of Refueling Stops (Hard)
// https://leetcode.com/problems/minimum-number-of-refueling-stops/
// Run: go run 0871_minimum_number_of_refueling_stops.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"os"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	h, reach, i, stops := &MaxHeap{}, startFuel, 0, 0
	for reach < target {
		for i < len(stations) && stations[i][0] <= reach { // stations we could have stopped at
			heap.Push(h, stations[i][1])
			i++
		}
		if h.Len() == 0 {
			return -1
		}
		reach += heap.Pop(h).(int) // retroactively stop at the biggest one
		stops++
	}
	return stops
}

func main() {
	assertEq(minRefuelStops(1, 1, [][]int{}), 0)
	assertEq(minRefuelStops(100, 1, [][]int{{10, 100}}), -1)
	assertEq(minRefuelStops(100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}), 2)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
