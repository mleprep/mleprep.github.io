//go:build ignore

// 402. Remove K Digits (Medium)
// https://leetcode.com/problems/remove-k-digits/
// Run: go run 0402_remove_k_digits.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func removeKdigits(num string, k int) string {
	st := []byte{} // increasing digits
	for i := 0; i < len(num); i++ {
		for k > 0 && len(st) > 0 && st[len(st)-1] > num[i] {
			st = st[:len(st)-1]
			k--
		}
		st = append(st, num[i])
	}
	st = st[:len(st)-k]
	res := strings.TrimLeft(string(st), "0")
	if res == "" {
		return "0"
	}
	return res
}

func main() {
	assertEq(removeKdigits("1432219", 3), "1219")
	assertEq(removeKdigits("10200", 1), "200")
	assertEq(removeKdigits("10", 2), "0")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
