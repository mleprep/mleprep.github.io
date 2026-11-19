//go:build ignore

// 190. Reverse Bits (Easy)
// https://leetcode.com/problems/reverse-bits/
// Run: go run 0190_reverse_bits.go lc.go
package main

import (
	"fmt"
	"os"
)

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res<<1 | num&1 // shift the lowest bit of num into res
		num >>= 1
	}
	return res
}

func main() {
	assertEq(reverseBits(43261596), 964176192)
	assertEq(reverseBits(4294967293), 3221225471)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
