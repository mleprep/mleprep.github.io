//go:build ignore

// 8. String to Integer (atoi) (Medium)
// https://leetcode.com/problems/string-to-integer-atoi/
// Run: go run 0008_string_to_integer_atoi.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func myAtoi(s string) int {
	i, n := 0, len(s)
	for i < n && s[i] == ' ' {
		i++
	}
	sign := 1
	if i < n && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	num := 0
	for ; i < n && s[i] >= '0' && s[i] <= '9'; i++ {
		num = num*10 + int(s[i]-'0')
		if num > math.MaxInt32+1 { // already out of range: stop before int64 overflows too
			break
		}
	}
	return max(math.MinInt32, min(math.MaxInt32, sign*num))
}

func main() {
	assertEq(myAtoi("42"), 42)
	assertEq(myAtoi("   -042"), -42)
	assertEq(myAtoi("1337c0d3"), 1337)
	assertEq(myAtoi("0-1"), 0)
	assertEq(myAtoi("-91283472332"), -2147483648)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
