//go:build ignore

// 6. Zigzag Conversion (Medium)
// https://leetcode.com/problems/zigzag-conversion/
// Run: go run 0006_zigzag_conversion.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	rows := make([][]byte, numRows)
	r, step := 0, 1
	for i := 0; i < len(s); i++ {
		rows[r] = append(rows[r], s[i])
		if r == 0 {
			step = 1
		} else if r == numRows-1 {
			step = -1
		}
		r += step
	}
	var sb strings.Builder
	for _, row := range rows {
		sb.Write(row)
	}
	return sb.String()
}

func main() {
	assertEq(convert("PAYPALISHIRING", 3), "PAHNAPLSIIGYIR")
	assertEq(convert("PAYPALISHIRING", 4), "PINALSIGYAHRPI")
	assertEq(convert("A", 1), "A")
	assertEq(convert("AB", 1), "AB")
	assertEq(convert("ABC", 2), "ACB")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
