//go:build ignore

// 881. Boats to Save People (Medium)
// https://leetcode.com/problems/boats-to-save-people/
// Run: go run 0881_boats_to_save_people.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	boats := 0
	for l, r := 0, len(people)-1; l <= r; r-- { // the heaviest always leaves
		if people[l]+people[r] <= limit {
			l++ // ...with the lightest if they fit
		}
		boats++
	}
	return boats
}

func main() {
	assertEq(numRescueBoats([]int{3, 2, 2, 1}, 3), 3)
	assertEq(numRescueBoats([]int{3, 5, 3, 4}, 5), 4)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
