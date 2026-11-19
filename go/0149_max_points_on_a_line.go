//go:build ignore

// 149. Max Points on a Line (Hard)
// https://leetcode.com/problems/max-points-on-a-line/
// Run: go run 0149_max_points_on_a_line.go lc.go
package main

import (
	"fmt"
	"os"
)

func maxPoints(points [][]int) int {
	best := 1
	for i, p := range points {
		slopes := map[[2]int]int{} // reduced direction -> points on that line through p
		for _, q := range points[i+1:] {
			dx, dy := q[0]-p[0], q[1]-p[1]
			g := gcd(abs(dx), abs(dy))
			dx, dy = dx/g, dy/g
			if dx < 0 || dx == 0 && dy < 0 { // one sign convention per direction
				dx, dy = -dx, -dy
			}
			key := [2]int{dx, dy}
			slopes[key]++
			best = max(best, slopes[key]+1)
		}
	}
	return best
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	assertEq(maxPoints([][]int{{1, 1}, {2, 2}, {3, 3}}), 3)
	assertEq(maxPoints([][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}), 4)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
