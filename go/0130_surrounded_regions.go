//go:build ignore

// 130. Surrounded Regions (Medium)
// https://leetcode.com/problems/surrounded-regions/
// Run: go run 0130_surrounded_regions.go lc.go
package main

import (
	"fmt"
	"os"
)

func solve(board [][]byte) {
	R, C := len(board), len(board[0])
	var mark func(r, c int) // "O"s connected to the border are safe
	mark = func(r, c int) {
		if r >= 0 && r < R && c >= 0 && c < C && board[r][c] == 'O' {
			board[r][c] = 'S'
			mark(r+1, c)
			mark(r-1, c)
			mark(r, c+1)
			mark(r, c-1)
		}
	}
	for r := 0; r < R; r++ {
		mark(r, 0)
		mark(r, C-1)
	}
	for c := 0; c < C; c++ {
		mark(0, c)
		mark(R-1, c)
	}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if board[r][c] == 'S' {
				board[r][c] = 'O'
			} else {
				board[r][c] = 'X'
			}
		}
	}
}

func main() {
	b := byteGrid("XXXX", "XOOX", "XXOX", "XOXX")
	solve(b)
	assertEq(b, byteGrid("XXXX", "XXXX", "XXXX", "XOXX"))
	b2 := byteGrid("X")
	solve(b2)
	assertEq(b2, byteGrid("X"))
	b3 := byteGrid("XXX", "XOX", "XXX")
	solve(b3)
	assertEq(b3, byteGrid("XXX", "XXX", "XXX"))
	b4 := byteGrid("OO", "OO")
	solve(b4)
	assertEq(b4, byteGrid("OO", "OO"))
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
