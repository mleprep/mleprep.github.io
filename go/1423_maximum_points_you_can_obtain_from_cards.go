//go:build ignore

// 1423. Maximum Points You Can Obtain from Cards (Medium)
// https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/
// Run: go run 1423_maximum_points_you_can_obtain_from_cards.go lc.go
package main

import (
	"fmt"
	"os"
)

func maxScore(cardPoints []int, k int) int {
	n, w := len(cardPoints), len(cardPoints)-k // minimise the middle window of size n-k
	total, s := 0, 0
	for i, x := range cardPoints {
		total += x
		if i < w {
			s += x
		}
	}
	best := s
	for i := w; i < n; i++ {
		s += cardPoints[i] - cardPoints[i-w]
		best = min(best, s)
	}
	return total - best
}

func main() {
	assertEq(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3), 12)
	assertEq(maxScore([]int{9, 7, 7, 9, 7, 7, 9}, 7), 55)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
