//go:build ignore

// 912. Sort an Array (Medium)
// https://leetcode.com/problems/sort-an-array/
// Run: go run 0912_sort_an_array.go lc.go
package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
)

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left, right := sortArray(nums[:mid:mid]), sortArray(nums[mid:])
	out := make([]int, 0, len(nums))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			out = append(out, left[i])
			i++
		} else {
			out = append(out, right[j])
			j++
		}
	}
	out = append(out, left[i:]...)
	return append(out, right[j:]...)
}

func main() {
	assertEq(sortArray([]int{5, 2, 3, 1}), []int{1, 2, 3, 5})
	assertEq(sortArray([]int{5, 1, 1, 2, 0, 0}), []int{0, 0, 1, 1, 2, 5})
	assertEq(sortArray([]int{1}), []int{1})
	rnd := rand.New(rand.NewSource(7))
	data := make([]int, 300)
	for i := range data {
		data[i] = rnd.Intn(101) - 50
	}
	want := append([]int{}, data...)
	sort.Ints(want)
	assertEq(sortArray(data), want)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
