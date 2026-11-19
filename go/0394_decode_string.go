//go:build ignore

// 394. Decode String (Medium)
// https://leetcode.com/problems/decode-string/
// Run: go run 0394_decode_string.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func decodeString(s string) string {
	type frame struct {
		prev  string
		times int
	}
	st, cur, k := []frame{}, "", 0
	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch {
		case ch >= '0' && ch <= '9':
			k = k*10 + int(ch-'0')
		case ch == '[':
			st = append(st, frame{cur, k})
			cur, k = "", 0
		case ch == ']':
			f := st[len(st)-1]
			st = st[:len(st)-1]
			cur = f.prev + strings.Repeat(cur, f.times)
		default:
			cur += string(ch)
		}
	}
	return cur
}

func main() {
	assertEq(decodeString("3[a]2[bc]"), "aaabcbc")
	assertEq(decodeString("3[a2[c]]"), "accaccacc")
	assertEq(decodeString("2[abc]3[cd]ef"), "abcabccdcdcdef")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
