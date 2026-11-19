//go:build ignore

// 85. Maximal Rectangle (Hard)
// https://leetcode.com/problems/maximal-rectangle/
// Run: go run 0085_maximal_rectangle.go lc.go
package main

import (
	"fmt"
	"os"
)

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	h := make([]int, len(matrix[0])) // histogram of 1s ending at this row
	best := 0
	for _, row := range matrix {
		for c, v := range row {
			if v == '1' {
				h[c]++
			} else {
				h[c] = 0
			}
		}
		st := []int{}
		for i := 0; i <= len(h); i++ {
			x := 0
			if i < len(h) {
				x = h[i]
			}
			for len(st) > 0 && h[st[len(st)-1]] >= x {
				top := h[st[len(st)-1]]
				st = st[:len(st)-1]
				left := -1
				if len(st) > 0 {
					left = st[len(st)-1]
				}
				best = max(best, top*(i-left-1))
			}
			st = append(st, i)
		}
	}
	return best
}

func main() {
	assertEq(maximalRectangle(byteGrid("10100", "10111", "11111", "10010")), 6)
	assertEq(maximalRectangle(byteGrid("0")), 0)
	assertEq(maximalRectangle(byteGrid("1")), 1)
	assertEq(maximalRectangle(byteGrid("0110", "1111", "1110")), 6)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
