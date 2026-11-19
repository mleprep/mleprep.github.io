//go:build ignore

// 1071. Greatest Common Divisor of Strings (Easy)
// https://leetcode.com/problems/greatest-common-divisor-of-strings/
// Run: go run 1071_greatest_common_divisor_of_strings.go lc.go
package main

import (
	"fmt"
	"os"
)

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 { // no common repeating unit
		return ""
	}
	a, b := len(str1), len(str2)
	for b != 0 {
		a, b = b, a%b
	}
	return str1[:a]
}

func main() {
	assertEq(gcdOfStrings("ABCABC", "ABC"), "ABC")
	assertEq(gcdOfStrings("ABABAB", "ABAB"), "AB")
	assertEq(gcdOfStrings("LEET", "CODE"), "")
	assertEq(gcdOfStrings("ABABABAB", "ABAB"), "ABAB")
	assertEq(gcdOfStrings("AAAA", "AA"), "AA")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
