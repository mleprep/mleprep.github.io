//go:build ignore

// 79. Word Search (Medium)
// https://leetcode.com/problems/word-search/
// Run: go run 0079_word_search.go lc.go
package main

import (
	"fmt"
	"os"
)

func exist(board [][]byte, word string) bool {
	R, C := len(board), len(board[0])
	var dfs func(r, c, i int) bool
	dfs = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}
		if r < 0 || r >= R || c < 0 || c >= C || board[r][c] != word[i] {
			return false
		}
		board[r][c] = '#' // mark as used on this path
		found := dfs(r+1, c, i+1) || dfs(r-1, c, i+1) || dfs(r, c+1, i+1) || dfs(r, c-1, i+1)
		board[r][c] = word[i]
		return found
	}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	B := byteGrid("ABCE", "SFCS", "ADEE")
	assertEq(exist(B, "ABCCED"), true)
	assertEq(exist(B, "SEE"), true)
	assertEq(exist(B, "ABCB"), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
