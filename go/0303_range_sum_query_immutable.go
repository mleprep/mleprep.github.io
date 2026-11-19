//go:build ignore

// 303. Range Sum Query - Immutable (Easy)
// https://leetcode.com/problems/range-sum-query-immutable/
// Run: go run 0303_range_sum_query_immutable.go lc.go
package main

import (
	"fmt"
	"os"
)

type NumArray struct {
	pre []int // pre[i] = sum of nums[:i]
}

func Constructor(nums []int) NumArray {
	pre := make([]int, len(nums)+1)
	for i, x := range nums {
		pre[i+1] = pre[i] + x
	}
	return NumArray{pre: pre}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.pre[right+1] - this.pre[left]
}

func main() {
	na := Constructor([]int{-2, 0, 3, -5, 2, -1})
	assertEq([]int{na.SumRange(0, 2), na.SumRange(2, 5), na.SumRange(0, 5)}, []int{1, -1, -3})
	one := Constructor([]int{5})
	assertEq(one.SumRange(0, 0), 5)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
