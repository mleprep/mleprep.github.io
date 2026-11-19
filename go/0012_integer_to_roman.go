//go:build ignore

// 12. Integer to Roman (Medium)
// https://leetcode.com/problems/integer-to-roman/
// Run: go run 0012_integer_to_roman.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func intToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var sb strings.Builder
	for i, v := range vals { // greedy: largest symbol first
		for num >= v {
			sb.WriteString(syms[i])
			num -= v
		}
	}
	return sb.String()
}

func main() {
	assertEq(intToRoman(3749), "MMMDCCXLIX")
	assertEq(intToRoman(58), "LVIII")
	assertEq(intToRoman(1994), "MCMXCIV")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
