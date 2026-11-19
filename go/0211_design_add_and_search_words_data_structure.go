//go:build ignore

// 211. Design Add and Search Words Data Structure (Medium)
// https://leetcode.com/problems/design-add-and-search-words-data-structure/
// Run: go run 0211_design_add_and_search_words_data_structure.go lc.go
package main

import (
	"fmt"
	"os"
)

type WordDictionary struct {
	children [26]*WordDictionary
	end      bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if node.children[c] == nil {
			node.children[c] = &WordDictionary{}
		}
		node = node.children[c]
	}
	node.end = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(node *WordDictionary, i int) bool
	dfs = func(node *WordDictionary, i int) bool {
		if i == len(word) {
			return node.end
		}
		if word[i] == '.' { // try every child
			for _, child := range node.children {
				if child != nil && dfs(child, i+1) {
					return true
				}
			}
			return false
		}
		child := node.children[word[i]-'a']
		return child != nil && dfs(child, i+1)
	}
	return dfs(this, 0)
}

func main() {
	wd := Constructor()
	for _, w := range []string{"bad", "dad", "mad"} {
		wd.AddWord(w)
	}
	assertEq([]bool{wd.Search("pad"), wd.Search("bad"), wd.Search(".ad"), wd.Search("b..")}, []bool{false, true, true, true})
	assertEq([]bool{wd.Search("..."), wd.Search("."), wd.Search("ba"), wd.Search("bad.")}, []bool{true, false, false, false})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
