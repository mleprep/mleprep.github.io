//go:build ignore

// 1871. Jump Game VII (Medium)
// https://leetcode.com/problems/jump-game-vii/
// Run: go run 1871_jump_game_vii.go lc.go
package main

import (
	"fmt"
	"os"
)

func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	ok := make([]bool, n) // ok[i]: index i is reachable
	ok[0] = true
	window := 0 // reachable indices in [i-maxJump, i-minJump]
	for i := 1; i < n; i++ {
		if i >= minJump && ok[i-minJump] {
			window++
		}
		if i > maxJump && ok[i-maxJump-1] {
			window--
		}
		ok[i] = window > 0 && s[i] == '0'
	}
	return ok[n-1]
}

func main() {
	assertEq(canReach("011010", 2, 3), true)
	assertEq(canReach("01101110", 2, 3), false)
	assertEq(canReach("00", 1, 1), true)
	assertEq(canReach("01", 1, 1), false)
	assertEq(canReach("0000000000", 8, 8), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
