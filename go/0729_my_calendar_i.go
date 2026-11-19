//go:build ignore

// 729. My Calendar I (Medium)
// https://leetcode.com/problems/my-calendar-i/
// Run: go run 0729_my_calendar_i.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

type MyCalendar struct {
	starts, ends []int // bookings sorted by start, never overlapping
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	i := sort.SearchInts(this.starts, start+1) // first booking starting after `start`
	if i > 0 && this.ends[i-1] > start {       // the previous one runs into us
		return false
	}
	if i < len(this.starts) && this.starts[i] < end { // we run into the next one
		return false
	}
	this.starts = insertAt(this.starts, i, start)
	this.ends = insertAt(this.ends, i, end)
	return true
}

func insertAt(s []int, i, v int) []int {
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = v
	return s
}

func main() {
	c := Constructor()
	assertEq([]bool{c.Book(10, 20), c.Book(15, 25), c.Book(20, 30)}, []bool{true, false, true})
	c2 := Constructor()
	assertEq([]bool{c2.Book(5, 10), c2.Book(10, 15), c2.Book(9, 11), c2.Book(0, 5), c2.Book(4, 6)}, []bool{true, true, false, true, false})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
