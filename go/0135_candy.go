//go:build ignore

// 135. Candy (Hard)
// https://leetcode.com/problems/candy/
// Run: go run 0135_candy.go lc.go
package main

import (
	"fmt"
	"os"
)

func candy(ratings []int) int {
	n := len(ratings)
	c := make([]int, n)
	for i := range c {
		c[i] = 1
	}
	for i := 1; i < n; i++ { // left neighbour rule
		if ratings[i] > ratings[i-1] {
			c[i] = c[i-1] + 1
		}
	}
	total := c[n-1]
	for i := n - 2; i >= 0; i-- { // right neighbour rule
		if ratings[i] > ratings[i+1] {
			c[i] = max(c[i], c[i+1]+1)
		}
		total += c[i]
	}
	return total
}

func main() {
	assertEq(candy([]int{1, 0, 2}), 5)
	assertEq(candy([]int{1, 2, 2}), 4)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
