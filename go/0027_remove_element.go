//go:build ignore

// 27. Remove Element (Easy)
// https://leetcode.com/problems/remove-element/
// Run: go run 0027_remove_element.go lc.go
package main

import (
	"fmt"
	"os"
)

func removeElement(nums []int, val int) int {
	k := 0
	for _, x := range nums {
		if x != val {
			nums[k] = x
			k++
		}
	}
	return k
}

func main() {
	a := []int{0, 1, 2, 2, 3, 0, 4, 2}
	k := removeElement(a, 2)
	assertUnordered(a[:k], []int{0, 0, 1, 3, 4})
	b := []int{3, 2, 2, 3}
	k = removeElement(b, 3)
	assertUnordered(b[:k], []int{2, 2})
	assertEq(removeElement([]int{1}, 1), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
