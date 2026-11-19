//go:build ignore

// 345. Reverse Vowels of a String (Easy)
// https://leetcode.com/problems/reverse-vowels-of-a-string/
// Run: go run 0345_reverse_vowels_of_a_string.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func reverseVowels(s string) string {
	vowel := func(c byte) bool { return strings.IndexByte("aeiouAEIOU", c) >= 0 }
	a := []byte(s)
	for i, j := 0, len(a)-1; i < j; {
		switch {
		case !vowel(a[i]):
			i++
		case !vowel(a[j]):
			j--
		default:
			a[i], a[j] = a[j], a[i]
			i, j = i+1, j-1
		}
	}
	return string(a)
}

func main() {
	assertEq(reverseVowels("IceCreAm"), "AceCreIm")
	assertEq(reverseVowels("leetcode"), "leotcede")
	assertEq(reverseVowels("bcd"), "bcd")
	assertEq(reverseVowels("aA"), "Aa")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
