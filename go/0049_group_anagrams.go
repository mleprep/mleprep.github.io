//go:build ignore

// 49. Group Anagrams (Medium)
// https://leetcode.com/problems/group-anagrams/
// Run: go run 0049_group_anagrams.go lc.go
package main

import (
	"fmt"
	"os"
)

func groupAnagrams(strs []string) [][]string {
	groups := map[[26]int][]string{} // letter counts are a valid map key
	for _, w := range strs {
		var key [26]int
		for _, ch := range w {
			key[ch-'a']++
		}
		groups[key] = append(groups[key], w)
	}
	res := make([][]string, 0, len(groups))
	for _, g := range groups {
		res = append(res, g)
	}
	return res
}

func main() {
	assertUnorderedDeep(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}), [][]string{{"ate", "eat", "tea"}, {"bat"}, {"nat", "tan"}})
	assertEq(groupAnagrams([]string{""}), [][]string{{""}})
	assertEq(groupAnagrams([]string{"a"}), [][]string{{"a"}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
