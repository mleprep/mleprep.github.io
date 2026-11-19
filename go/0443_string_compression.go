//go:build ignore

// 443. String Compression (Medium)
// https://leetcode.com/problems/string-compression/
// Run: go run 0443_string_compression.go lc.go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func compress(chars []byte) int {
	w := 0
	for i := 0; i < len(chars); {
		j := i
		for j < len(chars) && chars[j] == chars[i] {
			j++
		}
		chars[w] = chars[i]
		w++
		if j-i > 1 {
			for _, d := range strconv.Itoa(j - i) {
				chars[w] = byte(d)
				w++
			}
		}
		i = j
	}
	return w
}

func main() {
	a := []byte("aabbccc")
	k := compress(a)
	assertEq(string(a[:k]), "a2b2c3")
	b := []byte("a")
	k = compress(b)
	assertEq(string(b[:k]), "a")
	c := []byte("abbbbbbbbbbbb")
	k = compress(c)
	assertEq(string(c[:k]), "ab12")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
