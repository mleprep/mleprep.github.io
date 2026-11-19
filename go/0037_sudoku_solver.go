//go:build ignore

// 37. Sudoku Solver (Hard)
// https://leetcode.com/problems/sudoku-solver/
// Run: go run 0037_sudoku_solver.go lc.go
package main

import (
	"fmt"
	"os"
)

func solveSudoku(board [][]byte) {
	var rows, cols, boxes [9][10]bool
	empty := [][2]int{}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				empty = append(empty, [2]int{r, c})
			} else {
				v := board[r][c] - '0'
				rows[r][v], cols[c][v], boxes[r/3*3+c/3][v] = true, true, true
			}
		}
	}
	var bt func(k int) bool
	bt = func(k int) bool {
		if k == len(empty) {
			return true
		}
		r, c := empty[k][0], empty[k][1]
		b := r/3*3 + c/3
		for v := byte(1); v <= 9; v++ {
			if rows[r][v] || cols[c][v] || boxes[b][v] {
				continue
			}
			board[r][c] = '0' + v
			rows[r][v], cols[c][v], boxes[b][v] = true, true, true
			if bt(k + 1) {
				return true
			}
			rows[r][v], cols[c][v], boxes[b][v] = false, false, false
		}
		board[r][c] = '.'
		return false
	}
	bt(0)
}

func main() {
	board := byteGrid("53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1", "7...2...6", ".6....28.", "...419..5", "....8..79")
	solveSudoku(board)
	assertEq(board, byteGrid("534678912", "672195348", "198342567", "859761423", "426853791", "713924856", "961537284", "287419635", "345286179"))
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
