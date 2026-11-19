//go:build ignore

// 4. Median of Two Sorted Arrays (Hard)
// https://leetcode.com/problems/median-of-two-sorted-arrays/
// Run: go run 0004_median_of_two_sorted_arrays.go lc.go
package main

import (
	"fmt"
	"math"
	"os"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	A, B := nums1, nums2
	if len(A) > len(B) {
		A, B = B, A // binary search the shorter array
	}
	m, n := len(A), len(B)
	half := (m + n + 1) / 2
	lo, hi := 0, m
	for {
		i := (lo + hi) / 2 // i elements from A, j from B on the left side
		j := half - i
		al, ar, bl, br := math.MinInt, math.MaxInt, math.MinInt, math.MaxInt
		if i > 0 {
			al = A[i-1]
		}
		if i < m {
			ar = A[i]
		}
		if j > 0 {
			bl = B[j-1]
		}
		if j < n {
			br = B[j]
		}
		if al <= br && bl <= ar {
			if (m+n)%2 == 1 {
				return float64(max(al, bl))
			}
			return float64(max(al, bl)+min(ar, br)) / 2
		}
		if al > br {
			hi = i - 1
		} else {
			lo = i + 1
		}
	}
}

func main() {
	assertEq(findMedianSortedArrays([]int{1, 3}, []int{2}), 2.0)
	assertEq(findMedianSortedArrays([]int{1, 2}, []int{3, 4}), 2.5)
	assertEq(findMedianSortedArrays([]int{}, []int{1}), 1.0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
