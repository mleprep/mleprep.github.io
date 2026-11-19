//go:build ignore

// 451. Sort Characters By Frequency (Medium)
// https://leetcode.com/problems/sort-characters-by-frequency/
// Run: go run 0451_sort_characters_by_frequency.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func frequencySort(s string) string {
	cnt := map[byte]int{}
	for i := 0; i < len(s); i++ {
		cnt[s[i]]++
	}
	chars := make([]byte, 0, len(cnt))
	for c := range cnt {
		chars = append(chars, c)
	}
	sort.Slice(chars, func(i, j int) bool { return cnt[chars[i]] > cnt[chars[j]] })
	var sb strings.Builder
	for _, c := range chars {
		sb.WriteString(strings.Repeat(string(c), cnt[c]))
	}
	return sb.String()
}

func main() {
	r := frequencySort("tree")
	assertEq(r == "eert" || r == "eetr", true)
	r2 := frequencySort("Aabb")
	assertEq(r2 == "bbAa" || r2 == "bbaA", true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
