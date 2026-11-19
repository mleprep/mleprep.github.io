//go:build ignore

// 528. Random Pick with Weight (Medium)
// https://leetcode.com/problems/random-pick-with-weight/
// Run: go run 0528_random_pick_with_weight.go lc.go
package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
)

type Solution struct {
	pre []int // prefix sums of the weights
}

func Constructor(w []int) Solution {
	pre := make([]int, len(w))
	s := 0
	for i, x := range w {
		s += x
		pre[i] = s
	}
	return Solution{pre}
}

func (this *Solution) PickIndex() int {
	r := rand.Intn(this.pre[len(this.pre)-1]) // 0 <= r < total
	return sort.SearchInts(this.pre, r+1)     // first prefix sum > r
}

func main() {
	s := Constructor([]int{1, 3})
	picks := map[int]int{}
	for i := 0; i < 4000; i++ {
		picks[s.PickIndex()]++
	}
	assertEq(picks[0]+picks[1], 4000)
	assertEq(picks[1] > 2800 && picks[1] < 3200, true)
	s2 := Constructor([]int{5})
	assertEq(s2.PickIndex(), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
