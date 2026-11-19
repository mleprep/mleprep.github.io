//go:build ignore

// 1539. Kth Missing Positive Number (Easy)
// https://leetcode.com/problems/kth-missing-positive-number/
// Run: go run 1539_kth_missing_positive_number.go lc.go
package main

import (
	"fmt"
	"os"
)

func findKthPositive(arr []int, k int) int {
	lo, hi := 0, len(arr) // first index with arr[i]-(i+1) >= k missing numbers before it
	for lo < hi {
		mid := lo + (hi-lo)/2
		if arr[mid]-(mid+1) >= k {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo + k
}

func main() {
	assertEq(findKthPositive([]int{2, 3, 4, 7, 11}, 5), 9)
	assertEq(findKthPositive([]int{1, 2, 3, 4}, 2), 6)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
