//go:build ignore

// 2013. Detect Squares (Medium)
// https://leetcode.com/problems/detect-squares/
// Run: go run 2013_detect_squares.go lc.go
package main

import (
	"fmt"
	"os"
)

type DetectSquares struct {
	cnt map[[2]int]int // point -> how many times it was added
}

func Constructor() DetectSquares {
	return DetectSquares{cnt: map[[2]int]int{}}
}

func (this *DetectSquares) Add(point []int) {
	this.cnt[[2]int{point[0], point[1]}]++
}

func (this *DetectSquares) Count(point []int) int {
	x, y := point[0], point[1]
	total := 0
	for p, c := range this.cnt { // p is the corner diagonal to the query point
		px, py := p[0], p[1]
		if abs(px-x) != abs(py-y) || px == x {
			continue
		}
		total += c * this.cnt[[2]int{x, py}] * this.cnt[[2]int{px, y}]
	}
	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	ds := Constructor()
	for _, p := range [][]int{{3, 10}, {11, 2}, {3, 2}} {
		ds.Add(p)
	}
	assertEq(ds.Count([]int{11, 10}), 1)
	assertEq(ds.Count([]int{14, 8}), 0)
	ds.Add([]int{11, 2})
	assertEq(ds.Count([]int{11, 10}), 2)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
