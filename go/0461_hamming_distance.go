//go:build ignore

// 461. Hamming Distance (Easy)
// https://leetcode.com/problems/hamming-distance/
// Run: go run 0461_hamming_distance.go lc.go
package main

import (
	"fmt"
	"math/bits"
	"os"
)

func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y)) // the differing bits
}

func main() {
	assertEq(hammingDistance(1, 4), 2)
	assertEq(hammingDistance(3, 1), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
