//go:build ignore

// 208. Implement Trie (Prefix Tree) (Medium)
// https://leetcode.com/problems/implement-trie-prefix-tree/
// Run: go run 0208_implement_trie_prefix_tree.go lc.go
package main

import (
	"fmt"
	"os"
)

type Trie struct {
	children [26]*Trie
	end      bool // a word ends here
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if node.children[c] == nil {
			node.children[c] = &Trie{}
		}
		node = node.children[c]
	}
	node.end = true
}

func (this *Trie) walk(s string) *Trie {
	node := this
	for i := 0; i < len(s) && node != nil; i++ {
		node = node.children[s[i]-'a']
	}
	return node
}

func (this *Trie) Search(word string) bool {
	node := this.walk(word)
	return node != nil && node.end
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.walk(prefix) != nil
}

func main() {
	t := Constructor()
	t.Insert("apple")
	assertEq([]bool{t.Search("apple"), t.Search("app"), t.StartsWith("app")}, []bool{true, false, true})
	t.Insert("app")
	assertEq(t.Search("app"), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
