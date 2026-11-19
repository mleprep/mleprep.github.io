//go:build ignore

// 212. Word Search II (Hard)
// https://leetcode.com/problems/word-search-ii/
// Run: go run 0212_word_search_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

type trieNode struct {
	next [26]*trieNode
	word string // set at the node where a word ends
	kids int    // live children, for pruning
}

func findWords(board [][]byte, words []string) []string {
	root := &trieNode{}
	for _, w := range words {
		node := root
		for i := 0; i < len(w); i++ {
			c := w[i] - 'a'
			if node.next[c] == nil {
				node.next[c] = &trieNode{}
				node.kids++
			}
			node = node.next[c]
		}
		node.word = w
	}
	R, C, res := len(board), len(board[0]), []string{}
	var dfs func(r, c int, parent *trieNode)
	dfs = func(r, c int, parent *trieNode) {
		ch := board[r][c]
		node := parent.next[ch-'a']
		if node.word != "" {
			res = append(res, node.word)
			node.word = "" // report each word once
		}
		board[r][c] = '#'
		for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < R && nc >= 0 && nc < C && board[nr][nc] != '#' && node.next[board[nr][nc]-'a'] != nil {
				dfs(nr, nc, node)
			}
		}
		board[r][c] = ch
		if node.kids == 0 && node.word == "" { // prune an exhausted branch
			parent.next[ch-'a'] = nil
			parent.kids--
		}
	}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if root.next[board[r][c]-'a'] != nil {
				dfs(r, c, root)
			}
		}
	}
	return res
}

func main() {
	assertUnordered(findWords(byteGrid("oaan", "etae", "ihkr", "iflv"), []string{"oath", "pea", "eat", "rain"}), []string{"eat", "oath"})
	assertEq(findWords(byteGrid("ab", "cd"), []string{"abcb"}), []string{})
	assertEq(findWords(byteGrid("a"), []string{"a"}), []string{"a"})
	assertUnordered(findWords(byteGrid("ab", "cd"), []string{"ab", "cb", "ad", "bd", "ac", "ca", "da", "bc", "db", "adcb", "dabc", "abb", "acb"}), []string{"ab", "ac", "bd", "ca", "db"})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
