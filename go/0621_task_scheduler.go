//go:build ignore

// 621. Task Scheduler (Medium)
// https://leetcode.com/problems/task-scheduler/
// Run: go run 0621_task_scheduler.go lc.go
package main

import (
	"fmt"
	"os"
)

func leastInterval(tasks []byte, n int) int {
	var cnt [26]int
	for _, t := range tasks {
		cnt[t-'A']++
	}
	maxf, tied := 0, 0
	for _, c := range cnt {
		if c > maxf {
			maxf, tied = c, 1
		} else if c == maxf {
			tied++
		}
	}
	// (maxf-1) full frames of n+1 slots, plus the tasks tied for most frequent
	return max(len(tasks), (maxf-1)*(n+1)+tied)
}

func main() {
	assertEq(leastInterval([]byte("AAABBB"), 2), 8)
	assertEq(leastInterval([]byte("ACABDB"), 1), 6)
	assertEq(leastInterval([]byte("AAABBB"), 3), 10)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
