//go:build ignore

// 307. Range Sum Query - Mutable (Medium)
// https://leetcode.com/problems/range-sum-query-mutable/
// Run: go run 0307_range_sum_query_mutable.go lc.go
package main

import (
	"fmt"
	"os"
)

type NumArray struct {
	nums []int
	tree []int // Fenwick tree, 1-indexed
}

func Constructor(nums []int) NumArray {
	na := NumArray{nums: append([]int(nil), nums...), tree: make([]int, len(nums)+1)}
	for i, x := range nums {
		na.add(i+1, x)
	}
	return na
}

func (this *NumArray) add(i, v int) {
	for ; i < len(this.tree); i += i & -i {
		this.tree[i] += v
	}
}

func (this *NumArray) prefix(i int) int { // sum of nums[:i]
	s := 0
	for ; i > 0; i -= i & -i {
		s += this.tree[i]
	}
	return s
}

func (this *NumArray) Update(index int, val int) {
	this.add(index+1, val-this.nums[index])
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefix(right+1) - this.prefix(left)
}

func main() {
	na := Constructor([]int{1, 3, 5})
	assertEq(na.SumRange(0, 2), 9)
	na.Update(1, 2)
	assertEq([]int{na.SumRange(0, 2), na.SumRange(1, 1)}, []int{8, 2})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
