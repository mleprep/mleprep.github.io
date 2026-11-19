//go:build ignore

// 997. Find the Town Judge (Easy)
// https://leetcode.com/problems/find-the-town-judge/
// Run: go run 0997_find_the_town_judge.go lc.go
package main

import (
	"fmt"
	"os"
)

func findJudge(n int, trust [][]int) int {
	score := make([]int, n+1) // trusted-by count minus trusts count
	for _, t := range trust {
		score[t[0]]--
		score[t[1]]++
	}
	for p := 1; p <= n; p++ {
		if score[p] == n-1 {
			return p
		}
	}
	return -1
}

func main() {
	assertEq(findJudge(2, [][]int{{1, 2}}), 2)
	assertEq(findJudge(3, [][]int{{1, 3}, {2, 3}}), 3)
	assertEq(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}), -1)
	assertEq(findJudge(1, [][]int{}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
