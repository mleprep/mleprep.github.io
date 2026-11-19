//go:build ignore

// 406. Queue Reconstruction by Height (Medium)
// https://leetcode.com/problems/queue-reconstruction-by-height/
// Run: go run 0406_queue_reconstruction_by_height.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool { // tallest first; fewer people in front first
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	res := [][]int{}
	for _, p := range people { // everyone already placed is at least as tall: insert at index k
		k := p[1]
		res = append(res, p) // grow by one, then shift right from k
		copy(res[k+1:], res[k:])
		res[k] = p
	}
	return res
}

func main() {
	P := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	assertEq(reconstructQueue(P), [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}})
	assertEq(reconstructQueue([][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}), [][]int{{4, 0}, {5, 0}, {2, 2}, {3, 2}, {1, 4}, {6, 0}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
