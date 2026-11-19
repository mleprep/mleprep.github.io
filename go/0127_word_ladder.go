//go:build ignore

// 127. Word Ladder (Hard)
// https://leetcode.com/problems/word-ladder/
// Run: go run 0127_word_ladder.go lc.go
package main

import (
	"fmt"
	"os"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	words := map[string]bool{}
	for _, w := range wordList {
		words[w] = true
	}
	if !words[endWord] {
		return 0
	}
	delete(words, beginWord)
	q := []string{beginWord}
	for steps := 1; len(q) > 0; steps++ { // BFS, one level per step
		next := []string{}
		for _, w := range q {
			if w == endWord {
				return steps
			}
			b := []byte(w)
			for i := range b {
				orig := b[i]
				for ch := byte('a'); ch <= 'z'; ch++ {
					b[i] = ch
					if s := string(b); words[s] {
						delete(words, s) // visited
						next = append(next, s)
					}
				}
				b[i] = orig
			}
		}
		q = next
	}
	return 0
}

func main() {
	assertEq(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}), 5)
	assertEq(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
