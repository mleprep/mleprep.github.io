//go:build ignore

// 493. Reverse Pairs (Hard)
// https://leetcode.com/problems/reverse-pairs/
// Run: go run 0493_reverse_pairs.go lc.go
package main

import (
	"fmt"
	"os"
)

func reversePairs(nums []int) int {
	var sortCount func(a []int) ([]int, int) // merge sort that also counts the pairs
	sortCount = func(a []int) ([]int, int) {
		if len(a) <= 1 {
			return a, 0
		}
		mid := len(a) / 2
		left, cl := sortCount(append([]int(nil), a[:mid]...))
		right, cr := sortCount(append([]int(nil), a[mid:]...))
		count, j := cl+cr, 0
		for _, x := range left { // both halves sorted: j only moves forward
			for j < len(right) && x > 2*right[j] {
				j++
			}
			count += j
		}
		merged := make([]int, 0, len(a))
		i, k := 0, 0
		for i < len(left) && k < len(right) {
			if left[i] <= right[k] {
				merged = append(merged, left[i])
				i++
			} else {
				merged = append(merged, right[k])
				k++
			}
		}
		merged = append(append(merged, left[i:]...), right[k:]...)
		return merged, count
	}
	_, count := sortCount(nums)
	return count
}

func main() {
	assertEq(reversePairs([]int{1, 3, 2, 3, 1}), 2)
	assertEq(reversePairs([]int{2, 4, 3, 5, 1}), 3)
	assertEq(reversePairs([]int{2147483647, 2147483647, 2147483647}), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
