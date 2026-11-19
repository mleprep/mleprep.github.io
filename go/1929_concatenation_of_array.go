//go:build ignore

// 1929. Concatenation of Array (Easy)
// https://leetcode.com/problems/concatenation-of-array/
// Run: go run 1929_concatenation_of_array.go lc.go
package main

import (
	"fmt"
	"os"
)

func getConcatenation(nums []int) []int {
	return append(append([]int{}, nums...), nums...)
}

func main() {
	assertEq(getConcatenation([]int{1, 2, 1}), []int{1, 2, 1, 1, 2, 1})
	assertEq(getConcatenation([]int{1, 3, 2, 1}), []int{1, 3, 2, 1, 1, 3, 2, 1})
	assertEq(getConcatenation([]int{5}), []int{5, 5})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
