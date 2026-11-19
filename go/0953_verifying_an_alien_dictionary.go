//go:build ignore

// 953. Verifying an Alien Dictionary (Easy)
// https://leetcode.com/problems/verifying-an-alien-dictionary/
// Run: go run 0953_verifying_an_alien_dictionary.go lc.go
package main

import (
	"fmt"
	"os"
)

func isAlienSorted(words []string, order string) bool {
	var rank [26]int
	for i := 0; i < len(order); i++ {
		rank[order[i]-'a'] = i
	}
	less := func(a, b string) bool { // a <= b in the alien order
		for i := 0; i < len(a) && i < len(b); i++ {
			if a[i] != b[i] {
				return rank[a[i]-'a'] < rank[b[i]-'a']
			}
		}
		return len(a) <= len(b)
	}
	for i := 1; i < len(words); i++ {
		if !less(words[i-1], words[i]) {
			return false
		}
	}
	return true
}

func main() {
	assertEq(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"), true)
	assertEq(isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"), false)
	assertEq(isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
