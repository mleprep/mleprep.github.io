//go:build ignore

// 13. Roman to Integer (Easy)
// https://leetcode.com/problems/roman-to-integer/
// Run: go run 0013_roman_to_integer.go lc.go
package main

import (
	"fmt"
	"os"
)

func romanToInt(s string) int {
	val := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	total := 0
	for i := 0; i < len(s); i++ {
		v := val[s[i]]
		if i+1 < len(s) && val[s[i+1]] > v { // a smaller symbol before a larger one subtracts
			total -= v
		} else {
			total += v
		}
	}
	return total
}

func main() {
	assertEq(romanToInt("III"), 3)
	assertEq(romanToInt("LVIII"), 58)
	assertEq(romanToInt("MCMXCIV"), 1994)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
