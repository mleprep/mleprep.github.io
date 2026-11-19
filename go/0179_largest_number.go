//go:build ignore

// 179. Largest Number (Medium)
// https://leetcode.com/problems/largest-number/
// Run: go run 0179_largest_number.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	s := make([]string, len(nums))
	for i, x := range nums {
		s[i] = strconv.Itoa(x)
	}
	sort.Slice(s, func(i, j int) bool { return s[i]+s[j] > s[j]+s[i] })
	if s[0] == "0" {
		return "0"
	}
	return strings.Join(s, "")
}

func main() {
	assertEq(largestNumber([]int{10, 2}), "210")
	assertEq(largestNumber([]int{3, 30, 34, 5, 9}), "9534330")
	assertEq(largestNumber([]int{0, 0}), "0")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
