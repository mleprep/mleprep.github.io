//go:build ignore

// 51. N-Queens (Hard)
// https://leetcode.com/problems/n-queens/
// Run: go run 0051_n_queens.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func solveNQueens(n int) [][]string {
	res := [][]string{}
	cols, d1, d2 := map[int]bool{}, map[int]bool{}, map[int]bool{} // columns and both diagonals in use
	board := make([][]byte, n)
	for i := range board {
		board[i] = []byte(strings.Repeat(".", n))
	}
	var bt func(r int)
	bt = func(r int) {
		if r == n {
			rows := make([]string, n)
			for i, row := range board {
				rows[i] = string(row)
			}
			res = append(res, rows)
			return
		}
		for c := 0; c < n; c++ {
			if cols[c] || d1[r-c] || d2[r+c] {
				continue
			}
			cols[c], d1[r-c], d2[r+c] = true, true, true
			board[r][c] = 'Q'
			bt(r + 1)
			board[r][c] = '.'
			cols[c], d1[r-c], d2[r+c] = false, false, false
		}
	}
	bt(0)
	return res
}

func main() {
	assertUnordered(solveNQueens(4), [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}})
	assertEq(len(solveNQueens(8)), 92)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
