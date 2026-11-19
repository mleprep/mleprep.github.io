//go:build ignore

// 269. Alien Dictionary (Hard)
// https://leetcode.com/problems/alien-dictionary/
// Run: go run 0269_alien_dictionary.go lc.go
package main

import (
	"fmt"
	"os"
)

func alienOrder(words []string) string {
	adj := map[byte]map[byte]bool{}
	indeg := map[byte]int{}
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			if adj[w[i]] == nil {
				adj[w[i]] = map[byte]bool{}
				indeg[w[i]] = 0
			}
		}
	}
	for i := 0; i+1 < len(words); i++ {
		a, b := words[i], words[i+1]
		j := 0
		for j < len(a) && j < len(b) && a[j] == b[j] {
			j++
		}
		if j < len(a) && j < len(b) { // first difference: a[j] comes before b[j]
			if !adj[a[j]][b[j]] {
				adj[a[j]][b[j]] = true
				indeg[b[j]]++
			}
		} else if len(a) > len(b) {
			return "" // "abc" before "ab" is impossible
		}
	}
	q := []byte{} // Kahn's algorithm, letters taken alphabetically for a deterministic answer
	for c := byte('a'); c <= 'z'; c++ {
		if d, ok := indeg[c]; ok && d == 0 {
			q = append(q, c)
		}
	}
	out := []byte{}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		out = append(out, u)
		for v := byte('a'); v <= 'z'; v++ {
			if adj[u][v] {
				indeg[v]--
				if indeg[v] == 0 {
					q = append(q, v)
				}
			}
		}
	}
	if len(out) < len(indeg) {
		return "" // a cycle
	}
	return string(out)
}

func main() {
	assertEq(alienOrder([]string{"wrt", "wrf", "er", "ett", "rftt"}), "wertf")
	assertEq(alienOrder([]string{"z", "x", "z"}), "")
	assertEq(alienOrder([]string{"abc", "ab"}), "")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
