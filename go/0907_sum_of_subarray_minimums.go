//go:build ignore

// 907. Sum of Subarray Minimums (Medium)
// https://leetcode.com/problems/sum-of-subarray-minimums/
// Run: go run 0907_sum_of_subarray_minimums.go lc.go
package main

import (
	"fmt"
	"os"
)

func sumSubarrayMins(arr []int) int {
	const MOD = 1_000_000_007
	n := len(arr)
	left, right, st := make([]int, n), make([]int, n), []int{}
	for i := 0; i < n; i++ { // left[i]: choices of start where arr[i] is the min
		for len(st) > 0 && arr[st[len(st)-1]] > arr[i] {
			st = st[:len(st)-1]
		}
		left[i] = i + 1
		if len(st) > 0 {
			left[i] = i - st[len(st)-1]
		}
		st = append(st, i)
	}
	st = st[:0]
	for i := n - 1; i >= 0; i-- { // >= on this side so equal values are counted once
		for len(st) > 0 && arr[st[len(st)-1]] >= arr[i] {
			st = st[:len(st)-1]
		}
		right[i] = n - i
		if len(st) > 0 {
			right[i] = st[len(st)-1] - i
		}
		st = append(st, i)
	}
	sum := 0
	for i, a := range arr {
		sum = (sum + a*left[i]*right[i]) % MOD
	}
	return sum
}

func main() {
	assertEq(sumSubarrayMins([]int{3, 1, 2, 4}), 17)
	assertEq(sumSubarrayMins([]int{11, 81, 94, 43, 3}), 444)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
