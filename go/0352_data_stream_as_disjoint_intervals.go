//go:build ignore

// 352. Data Stream as Disjoint Intervals (Hard)
// https://leetcode.com/problems/data-stream-as-disjoint-intervals/
// Run: go run 0352_data_stream_as_disjoint_intervals.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

type SummaryRanges struct {
	iv [][]int // sorted, disjoint [start, end]
}

func Constructor() SummaryRanges {
	return SummaryRanges{}
}

func (this *SummaryRanges) AddNum(value int) {
	iv := this.iv
	i := sort.Search(len(iv), func(k int) bool { return iv[k][0] > value }) // first interval starting after value
	if i > 0 && iv[i-1][1] >= value {
		return // already covered
	}
	left := i > 0 && iv[i-1][1] == value-1
	right := i < len(iv) && iv[i][0] == value+1
	switch {
	case left && right: // value bridges two intervals
		iv[i-1][1] = iv[i][1]
		this.iv = append(iv[:i], iv[i+1:]...)
	case left:
		iv[i-1][1] = value
	case right:
		iv[i][0] = value
	default:
		iv = append(iv, []int{value, value})
		copy(iv[i+1:], iv[i:])
		iv[i] = []int{value, value}
		this.iv = iv
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	out := make([][]int, len(this.iv))
	for i, x := range this.iv {
		out[i] = []int{x[0], x[1]}
	}
	return out
}

func main() {
	sr := Constructor()
	out := [][][]int{}
	for _, v := range []int{1, 3, 7, 2, 6} {
		sr.AddNum(v)
		out = append(out, sr.GetIntervals())
	}
	assertEq(out, [][][]int{{{1, 1}}, {{1, 1}, {3, 3}}, {{1, 1}, {3, 3}, {7, 7}}, {{1, 3}, {7, 7}}, {{1, 3}, {6, 7}}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
