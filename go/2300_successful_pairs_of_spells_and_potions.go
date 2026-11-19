//go:build ignore

// 2300. Successful Pairs of Spells and Potions (Medium)
// https://leetcode.com/problems/successful-pairs-of-spells-and-potions/
// Run: go run 2300_successful_pairs_of_spells_and_potions.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	m := len(potions)
	res := make([]int, len(spells))
	for i, s := range spells {
		need := (int(success) + s - 1) / s // weakest potion that works
		res[i] = m - sort.SearchInts(potions, need)
	}
	return res
}

func main() {
	assertEq(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7), []int{4, 0, 3})
	assertEq(successfulPairs([]int{3, 1, 2}, []int{8, 5, 8}, 16), []int{2, 0, 2})
	assertEq(successfulPairs([]int{100000}, []int{100000}, 10000000000), []int{1})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
