//go:build ignore

// 253. Meeting Rooms II (Medium)
// https://leetcode.com/problems/meeting-rooms-ii/
// Run: go run 0253_meeting_rooms_ii.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	starts, ends := make([]int, n), make([]int, n)
	for i, iv := range intervals {
		starts[i], ends[i] = iv[0], iv[1]
	}
	sort.Ints(starts)
	sort.Ints(ends)
	rooms, best, j := 0, 0, 0
	for _, s := range starts { // sweep: free the rooms whose meetings ended by s
		for ends[j] <= s {
			rooms--
			j++
		}
		rooms++
		best = max(best, rooms)
	}
	return best
}

func main() {
	assertEq(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}), 2)
	assertEq(minMeetingRooms([][]int{{7, 10}, {2, 4}}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
