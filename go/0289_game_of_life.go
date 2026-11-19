//go:build ignore

// 289. Game of Life (Medium)
// https://leetcode.com/problems/game-of-life/
// Run: go run 0289_game_of_life.go lc.go
package main

import (
	"fmt"
	"os"
)

func gameOfLife(board [][]int) {
	R, C := len(board), len(board[0])
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			live := 0
			for i := max(0, r-1); i < min(R, r+2); i++ {
				for j := max(0, c-1); j < min(C, c+2); j++ {
					live += board[i][j] & 1 // bit 0 = current state
				}
			}
			live -= board[r][c] & 1
			if live == 3 || (live == 2 && board[r][c]&1 == 1) {
				board[r][c] |= 2 // bit 1 = next state
			}
		}
	}
	for r := range board {
		for c := range board[r] {
			board[r][c] >>= 1
		}
	}
}

func main() {
	b := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(b)
	assertEq(b, [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}})
	b2 := [][]int{{1, 1}, {1, 0}}
	gameOfLife(b2)
	assertEq(b2, [][]int{{1, 1}, {1, 1}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
