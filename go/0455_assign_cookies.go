//go:build ignore

// 455. Assign Cookies (Easy)
// https://leetcode.com/problems/assign-cookies/
// Run: go run 0455_assign_cookies.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	child := 0
	for _, size := range s { // each cookie goes to the least greedy child it satisfies
		if child < len(g) && size >= g[child] {
			child++
		}
	}
	return child
}

func main() {
	assertEq(findContentChildren([]int{1, 2, 3}, []int{1, 1}), 1)
	assertEq(findContentChildren([]int{1, 2}, []int{1, 2, 3}), 2)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
