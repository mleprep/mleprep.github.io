//go:build ignore

// 649. Dota2 Senate (Medium)
// https://leetcode.com/problems/dota2-senate/
// Run: go run 0649_dota2_senate.go lc.go
package main

import (
	"fmt"
	"os"
)

func predictPartyVictory(senate string) string {
	n := len(senate)
	r, d := []int{}, []int{} // turn indices of each party's active senators
	for i := 0; i < n; i++ {
		if senate[i] == 'R' {
			r = append(r, i)
		} else {
			d = append(d, i)
		}
	}
	for len(r) > 0 && len(d) > 0 {
		a, b := r[0], d[0]
		r, d = r[1:], d[1:]
		if a < b { // the earlier senator bans the other and votes again next round
			r = append(r, a+n)
		} else {
			d = append(d, b+n)
		}
	}
	if len(r) > 0 {
		return "Radiant"
	}
	return "Dire"
}

func main() {
	assertEq(predictPartyVictory("RD"), "Radiant")
	assertEq(predictPartyVictory("RDD"), "Dire")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
