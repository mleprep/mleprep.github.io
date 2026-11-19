//go:build ignore

// 168. Excel Sheet Column Title (Easy)
// https://leetcode.com/problems/excel-sheet-column-title/
// Run: go run 0168_excel_sheet_column_title.go lc.go
package main

import (
	"fmt"
	"os"
)

func convertToTitle(columnNumber int) string {
	out := []byte{}
	for columnNumber > 0 { // bijective base 26: no zero digit, so shift by one first
		columnNumber--
		out = append(out, byte('A'+columnNumber%26))
		columnNumber /= 26
	}
	for l, r := 0, len(out)-1; l < r; l, r = l+1, r-1 {
		out[l], out[r] = out[r], out[l]
	}
	return string(out)
}

func main() {
	assertEq(convertToTitle(1), "A")
	assertEq(convertToTitle(28), "AB")
	assertEq(convertToTitle(701), "ZY")
	assertEq(convertToTitle(26), "Z")
	assertEq(convertToTitle(52), "AZ")
	assertEq(convertToTitle(2147483647), "FXSHRXW")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
