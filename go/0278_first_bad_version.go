//go:build ignore

// 278. First Bad Version (Easy)
// https://leetcode.com/problems/first-bad-version/
// Run: go run 0278_first_bad_version.go lc.go
package main

import (
	"fmt"
	"os"
)

var bad int // provided by LeetCode (the tests set it)

func isBadVersion(version int) bool { // provided by LeetCode
	return version >= bad
}

func firstBadVersion(n int) int {
	lo, hi := 1, n
	for lo < hi {
		mid := lo + (hi-lo)/2
		if isBadVersion(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	bad = 4
	assertEq(firstBadVersion(5), 4)
	bad = 1
	assertEq(firstBadVersion(1), 1)
	bad = 1702766719
	assertEq(firstBadVersion(2126753390), 1702766719)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
