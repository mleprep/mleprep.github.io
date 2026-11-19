//go:build ignore

// 904. Fruit Into Baskets (Medium)
// https://leetcode.com/problems/fruit-into-baskets/
// Run: go run 0904_fruit_into_baskets.go lc.go
package main

import (
	"fmt"
	"os"
)

func totalFruit(fruits []int) int {
	cnt := map[int]int{}
	l, best := 0, 0
	for r, f := range fruits {
		cnt[f]++
		for len(cnt) > 2 {
			cnt[fruits[l]]--
			if cnt[fruits[l]] == 0 {
				delete(cnt, fruits[l])
			}
			l++
		}
		best = max(best, r-l+1)
	}
	return best
}

func main() {
	assertEq(totalFruit([]int{1, 2, 1}), 3)
	assertEq(totalFruit([]int{0, 1, 2, 2}), 3)
	assertEq(totalFruit([]int{1, 2, 3, 2, 2}), 4)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
