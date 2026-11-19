//go:build ignore

// 36. Valid Sudoku (Medium)
// https://leetcode.com/problems/valid-sudoku/
// Run: go run 0036_valid_sudoku.go lc.go
package main

import (
	"fmt"
	"os"
)

func isValidSudoku(board [][]byte) bool {
	var rows, cols, boxes [9][9]bool
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			v, b := board[r][c]-'1', r/3*3+c/3
			if rows[r][v] || cols[c][v] || boxes[b][v] {
				return false
			}
			rows[r][v], cols[c][v], boxes[b][v] = true, true, true
		}
	}
	return true
}

func main() {
	board := byteGrid("53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1",
		"7...2...6", ".6....28.", "...419..5", "....8..79")
	assertEq(isValidSudoku(board), true)
	board[0][0] = '8' // column 0 already has an 8
	assertEq(isValidSudoku(board), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
