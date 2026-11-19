//go:build ignore

// 17. Letter Combinations of a Phone Number (Medium)
// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
// Run: go run 0017_letter_combinations_of_a_phone_number.go lc.go
package main

import (
	"fmt"
	"os"
)

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	pad := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res := []string{}
	var bt func(i int, cur string)
	bt = func(i int, cur string) {
		if i == len(digits) {
			res = append(res, cur)
			return
		}
		for _, ch := range pad[digits[i]-'0'] {
			bt(i+1, cur+string(ch))
		}
	}
	bt(0, "")
	return res
}

func main() {
	assertUnordered(letterCombinations("23"), []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"})
	assertEq(letterCombinations(""), []string{})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
