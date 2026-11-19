//go:build ignore

// 374. Guess Number Higher or Lower (Easy)
// https://leetcode.com/problems/guess-number-higher-or-lower/
// Run: go run 0374_guess_number_higher_or_lower.go lc.go
package main

import (
	"fmt"
	"os"
)

var pick int // provided by LeetCode (the tests set it)

func guess(num int) int { // provided by LeetCode
	if num == pick {
		return 0
	}
	if pick < num {
		return -1
	}
	return 1
}

func guessNumber(n int) int {
	lo, hi := 1, n
	for {
		mid := lo + (hi-lo)/2
		switch guess(mid) {
		case 0:
			return mid
		case -1:
			hi = mid - 1
		default:
			lo = mid + 1
		}
	}
}

func main() {
	for _, tc := range [][2]int{{10, 6}, {1, 1}, {2, 1}, {2, 2}, {2126753390, 1702766719}} {
		pick = tc[1]
		assertEq(guessNumber(tc[0]), tc[1])
	}
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
