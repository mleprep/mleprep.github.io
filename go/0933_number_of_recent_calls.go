//go:build ignore

// 933. Number of Recent Calls (Easy)
// https://leetcode.com/problems/number-of-recent-calls/
// Run: go run 0933_number_of_recent_calls.go lc.go
package main

import (
	"fmt"
	"os"
)

type RecentCounter struct {
	q []int // ping times within the last 3000 ms
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.q = append(this.q, t)
	for this.q[0] < t-3000 {
		this.q = this.q[1:]
	}
	return len(this.q)
}

func main() {
	rc := Constructor()
	assertEq([]int{rc.Ping(1), rc.Ping(100), rc.Ping(3001), rc.Ping(3002)}, []int{1, 2, 3, 3})
	rc2 := Constructor()
	assertEq([]int{rc2.Ping(1), rc2.Ping(3001), rc2.Ping(3002), rc2.Ping(6003)}, []int{1, 2, 2, 1})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
