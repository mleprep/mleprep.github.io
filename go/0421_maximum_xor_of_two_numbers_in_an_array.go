//go:build ignore

// 421. Maximum XOR of Two Numbers in an Array (Medium)
// https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/
// Run: go run 0421_maximum_xor_of_two_numbers_in_an_array.go lc.go
package main

import (
	"fmt"
	"os"
)

func findMaximumXOR(nums []int) int {
	type node struct{ next [2]*node }
	root := &node{}
	for _, x := range nums { // binary trie of the numbers, high bit first
		n := root
		for b := 30; b >= 0; b-- {
			bit := x >> b & 1
			if n.next[bit] == nil {
				n.next[bit] = &node{}
			}
			n = n.next[bit]
		}
	}
	best := 0
	for _, x := range nums {
		n, cur := root, 0
		for b := 30; b >= 0; b-- {
			bit := x >> b & 1
			if n.next[1-bit] != nil { // take the opposite bit whenever possible
				cur |= 1 << b
				n = n.next[1-bit]
			} else {
				n = n.next[bit]
			}
		}
		best = max(best, cur)
	}
	return best
}

func main() {
	assertEq(findMaximumXOR([]int{3, 10, 5, 25, 2, 8}), 28)
	assertEq(findMaximumXOR([]int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}), 127)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
