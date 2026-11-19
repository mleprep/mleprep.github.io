//go:build ignore

// 648. Replace Words (Medium)
// https://leetcode.com/problems/replace-words/
// Run: go run 0648_replace_words.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	type node struct {
		next [26]*node
		end  bool
	}
	root := &node{}
	for _, w := range dictionary {
		n := root
		for i := 0; i < len(w); i++ {
			c := w[i] - 'a'
			if n.next[c] == nil {
				n.next[c] = &node{}
			}
			n = n.next[c]
		}
		n.end = true
	}
	shortest := func(word string) string { // the shortest root that prefixes word
		n := root
		for i := 0; i < len(word); i++ {
			if n = n.next[word[i]-'a']; n == nil {
				return word
			}
			if n.end {
				return word[:i+1]
			}
		}
		return word
	}
	words := strings.Fields(sentence)
	for i, w := range words {
		words[i] = shortest(w)
	}
	return strings.Join(words, " ")
}

func main() {
	assertEq(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"), "the cat was rat by the bat")
	assertEq(replaceWords([]string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs"), "a a b c")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
