//go:build ignore

// 252. Meeting Rooms (Easy)
// https://leetcode.com/problems/meeting-rooms/
// Run: go run 0252_meeting_rooms.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] { // starts before the previous one ends
			return false
		}
	}
	return true
}

func main() {
	assertEq(canAttendMeetings([][]int{{0, 30}, {5, 10}, {15, 20}}), false)
	assertEq(canAttendMeetings([][]int{{7, 10}, {2, 4}}), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
