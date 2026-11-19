//go:build ignore

// 784. Letter Case Permutation (Medium)
// https://leetcode.com/problems/letter-case-permutation/
// Run: go run 0784_letter_case_permutation.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func letterCasePermutation(s string) []string {
	res := []string{""}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		next := []string{}
		for _, p := range res {
			if unicode.IsLetter(rune(ch)) { // branch: lower then upper
				next = append(next, p+strings.ToLower(string(ch)), p+strings.ToUpper(string(ch)))
			} else {
				next = append(next, p+string(ch))
			}
		}
		res = next
	}
	return res
}

func main() {
	assertUnordered(letterCasePermutation("a1b2"), []string{"A1B2", "A1b2", "a1B2", "a1b2"})
	assertEq(letterCasePermutation("3z4"), []string{"3z4", "3Z4"})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
