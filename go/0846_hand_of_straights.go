//go:build ignore

// 846. Hand of Straights (Medium)
// https://leetcode.com/problems/hand-of-straights/
// Run: go run 0846_hand_of_straights.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	cnt := map[int]int{}
	for _, x := range hand {
		cnt[x]++
	}
	keys := make([]int, 0, len(cnt))
	for x := range cnt {
		keys = append(keys, x)
	}
	sort.Ints(keys)
	for _, x := range keys { // the smallest card left must start a group
		c := cnt[x]
		if c == 0 {
			continue
		}
		for y := x; y < x+groupSize; y++ {
			if cnt[y] < c {
				return false
			}
			cnt[y] -= c
		}
	}
	return true
}

func main() {
	assertEq(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3), true)
	assertEq(isNStraightHand([]int{1, 2, 3, 4, 5}, 4), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
