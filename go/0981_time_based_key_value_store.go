//go:build ignore

// 981. Time Based Key-Value Store (Medium)
// https://leetcode.com/problems/time-based-key-value-store/
// Run: go run 0981_time_based_key_value_store.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

type entry struct {
	ts    int
	value string
}

type TimeMap struct {
	data map[string][]entry // timestamps arrive increasing
}

func Constructor() TimeMap {
	return TimeMap{data: map[string][]entry{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.data[key] = append(this.data[key], entry{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	arr := this.data[key]
	i := sort.Search(len(arr), func(i int) bool { return arr[i].ts > timestamp }) // first ts > timestamp
	if i == 0 {
		return ""
	}
	return arr[i-1].value
}

func main() {
	tm := Constructor()
	tm.Set("foo", "bar", 1)
	assertEq(tm.Get("foo", 1), "bar")
	assertEq(tm.Get("foo", 3), "bar")
	tm.Set("foo", "bar2", 4)
	assertEq([]string{tm.Get("foo", 4), tm.Get("foo", 5), tm.Get("foo", 0)}, []string{"bar2", "bar2", ""})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
