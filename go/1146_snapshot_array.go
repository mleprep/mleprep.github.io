//go:build ignore

// 1146. Snapshot Array (Medium)
// https://leetcode.com/problems/snapshot-array/
// Run: go run 1146_snapshot_array.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

type SnapshotArray struct {
	hist   [][][2]int // per index: (snap id, value) changes, in snap order
	snapID int
}

func Constructor(length int) SnapshotArray {
	hist := make([][][2]int, length)
	for i := range hist {
		hist[i] = [][2]int{{-1, 0}} // everything starts at 0
	}
	return SnapshotArray{hist: hist}
}

func (this *SnapshotArray) Set(index int, val int) {
	h := this.hist[index]
	if h[len(h)-1][0] == this.snapID { // overwrite within the same snapshot
		h[len(h)-1][1] = val
	} else {
		this.hist[index] = append(h, [2]int{this.snapID, val})
	}
}

func (this *SnapshotArray) Snap() int {
	this.snapID++
	return this.snapID - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	h := this.hist[index]
	i := sort.Search(len(h), func(k int) bool { return h[k][0] > snap_id }) - 1 // last change at or before snap_id
	return h[i][1]
}

func main() {
	sa := Constructor(3)
	sa.Set(0, 5)
	assertEq(sa.Snap(), 0)
	sa.Set(0, 6)
	assertEq([]int{sa.Get(0, 0), sa.Get(0, 1), sa.Get(1, 0)}, []int{5, 6, 0})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
