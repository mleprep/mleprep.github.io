//go:build ignore

// 43. Multiply Strings (Medium)
// https://leetcode.com/problems/multiply-strings/
// Run: go run 0043_multiply_strings.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	pos := make([]int, m+n) // digit i*j lands in pos[i+j+1], its carry in pos[i+j]
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			s := int(num1[i]-'0')*int(num2[j]-'0') + pos[i+j+1]
			pos[i+j+1] = s % 10
			pos[i+j] += s / 10
		}
	}
	var sb strings.Builder
	for _, d := range pos {
		if sb.Len() > 0 || d != 0 { // skip leading zeros
			sb.WriteByte(byte('0' + d))
		}
	}
	if sb.Len() == 0 {
		return "0"
	}
	return sb.String()
}

func main() {
	assertEq(multiply("2", "3"), "6")
	assertEq(multiply("123", "456"), "56088")
	assertEq(multiply("0", "52"), "0")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
