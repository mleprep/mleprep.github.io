//go:build ignore

// 229. Majority Element II (Medium)
// https://leetcode.com/problems/majority-element-ii/
// Run: go run 0229_majority_element_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func majorityElement(nums []int) []int {
	c1, c2, n1, n2 := 0, 1, 0, 0 // two candidates (distinct initial values)
	for _, x := range nums {
		switch {
		case x == c1:
			n1++
		case x == c2:
			n2++
		case n1 == 0:
			c1, n1 = x, 1
		case n2 == 0:
			c2, n2 = x, 1
		default:
			n1, n2 = n1-1, n2-1
		}
	}
	res := []int{}
	for _, c := range []int{c1, c2} {
		cnt := 0
		for _, x := range nums {
			if x == c {
				cnt++
			}
		}
		if cnt > len(nums)/3 {
			res = append(res, c)
		}
	}
	return res
}

func main() {
	assertEq(majorityElement([]int{3, 2, 3}), []int{3})
	assertEq(majorityElement([]int{1}), []int{1})
	assertUnordered(majorityElement([]int{1, 2}), []int{1, 2})
	assertUnordered(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}), []int{1, 2})
	assertEq(majorityElement([]int{1, 2, 3, 4}), []int{})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
