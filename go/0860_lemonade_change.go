//go:build ignore

// 860. Lemonade Change (Easy)
// https://leetcode.com/problems/lemonade-change/
// Run: go run 0860_lemonade_change.go lc.go
package main

import (
	"fmt"
	"os"
)

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, b := range bills {
		switch {
		case b == 5:
			five++
		case b == 10:
			five, ten = five-1, ten+1
		case ten > 0: // give 10 + 5 for a 20: keep fives, they are more useful
			five, ten = five-1, ten-1
		default:
			five -= 3
		}
		if five < 0 {
			return false
		}
	}
	return true
}

func main() {
	assertEq(lemonadeChange([]int{5, 5, 5, 10, 20}), true)
	assertEq(lemonadeChange([]int{5, 5, 10, 10, 20}), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
