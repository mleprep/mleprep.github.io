//go:build ignore

// 91. Decode Ways (Medium)
// https://leetcode.com/problems/decode-ways/
// Run: go run 0091_decode_ways.go lc.go
package main

import (
	"fmt"
	"os"
)

func numDecodings(s string) int {
	a, b := 1, 0 // dp[i-2], dp[i-1]
	if s[0] != '0' {
		b = 1
	}
	for i := 2; i <= len(s); i++ {
		cur := 0
		if s[i-1] != '0' {
			cur = b // last digit alone
		}
		if two := int(s[i-2]-'0')*10 + int(s[i-1]-'0'); two >= 10 && two <= 26 {
			cur += a // last two digits together
		}
		a, b = b, cur
	}
	return b
}

func main() {
	assertEq(numDecodings("12"), 2)
	assertEq(numDecodings("226"), 3)
	assertEq(numDecodings("06"), 0)
	assertEq(numDecodings("10"), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
