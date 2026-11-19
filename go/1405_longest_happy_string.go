//go:build ignore

// 1405. Longest Happy String (Medium)
// https://leetcode.com/problems/longest-happy-string/
// Run: go run 1405_longest_happy_string.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func longestDiverseString(a int, b int, c int) string {
	cnt := []int{a, b, c}
	out := []byte{}
	for {
		order := []int{0, 1, 2} // letters by remaining count, most first
		sort.Slice(order, func(i, j int) bool { return cnt[order[i]] > cnt[order[j]] })
		placed := false
		for _, x := range order {
			ch := byte('a' + x)
			n := len(out)
			if cnt[x] == 0 || (n >= 2 && out[n-1] == ch && out[n-2] == ch) {
				continue // would make three in a row
			}
			out = append(out, ch)
			cnt[x]--
			placed = true
			break
		}
		if !placed {
			return string(out)
		}
	}
}

func main() {
	happy := func(s string, a, b, c int) bool { // any valid answer of maximum length is accepted
		v := []int{a, b, c}
		sort.Sort(sort.Reverse(sort.IntSlice(v)))
		x, y, z := v[0], v[1], v[2]
		return !strings.Contains(s, "aaa") && !strings.Contains(s, "bbb") && !strings.Contains(s, "ccc") &&
			strings.Count(s, "a") <= a && strings.Count(s, "b") <= b && strings.Count(s, "c") <= c &&
			len(s) == y+z+min(x, 2*(y+z+1))
	}
	for _, tc := range [][3]int{{1, 1, 7}, {7, 1, 0}, {2, 2, 1}, {0, 0, 0}, {0, 8, 11}, {5, 5, 5}, {0, 0, 9}} {
		assertEq(happy(longestDiverseString(tc[0], tc[1], tc[2]), tc[0], tc[1], tc[2]), true)
	}
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
