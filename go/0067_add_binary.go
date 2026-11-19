//go:build ignore

// 67. Add Binary (Easy)
// https://leetcode.com/problems/add-binary/
// Run: go run 0067_add_binary.go lc.go
package main

import (
	"fmt"
	"os"
)

func addBinary(a string, b string) string {
	out := []byte{}
	for i, j, carry := len(a)-1, len(b)-1, 0; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		s := carry
		if i >= 0 {
			s += int(a[i] - '0')
		}
		if j >= 0 {
			s += int(b[j] - '0')
		}
		out = append(out, byte('0'+s%2))
		carry = s / 2
	}
	for l, r := 0, len(out)-1; l < r; l, r = l+1, r-1 {
		out[l], out[r] = out[r], out[l]
	}
	return string(out)
}

func main() {
	assertEq(addBinary("11", "1"), "100")
	assertEq(addBinary("1010", "1011"), "10101")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
