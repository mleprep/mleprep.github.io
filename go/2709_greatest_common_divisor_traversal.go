//go:build ignore

// 2709. Greatest Common Divisor Traversal (Hard)
// https://leetcode.com/problems/greatest-common-divisor-traversal/
// Run: go run 2709_greatest_common_divisor_traversal.go lc.go
package main

import (
	"fmt"
	"os"
)

func canTraverseAllPairs(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}
	M := 0
	for _, x := range nums {
		if x == 1 {
			return false // 1 shares no factor with anything
		}
		M = max(M, x)
	}
	spf := make([]int, M+1) // smallest prime factor
	for i := range spf {
		spf[i] = i
	}
	for i := 2; i*i <= M; i++ {
		if spf[i] == i {
			for j := i * i; j <= M; j += i {
				if spf[j] == j {
					spf[j] = i
				}
			}
		}
	}
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	find := func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	first := map[int]int{} // prime -> first index divisible by it
	for i, x := range nums {
		for x > 1 {
			p := spf[x]
			if j, ok := first[p]; ok {
				parent[find(i)] = find(j)
			} else {
				first[p] = i
			}
			for x%p == 0 {
				x /= p
			}
		}
	}
	for i := 1; i < n; i++ {
		if find(i) != find(0) {
			return false
		}
	}
	return true
}

func main() {
	assertEq(canTraverseAllPairs([]int{2, 3, 6}), true)
	assertEq(canTraverseAllPairs([]int{3, 9, 5}), false)
	assertEq(canTraverseAllPairs([]int{4, 3, 12, 8}), true)
	assertEq(canTraverseAllPairs([]int{1}), true)
	assertEq(canTraverseAllPairs([]int{1, 1}), false)
	assertEq(canTraverseAllPairs([]int{2, 2}), true)
	assertEq(canTraverseAllPairs([]int{35, 77, 22, 6}), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
