//go:build ignore

// 55. Jump Game (Medium)
// https://leetcode.com/problems/jump-game/
// Run: go run 0055_jump_game.go lc.go
package main

import (
	"fmt"
	"os"
)

func canJump(nums []int) bool {
	reach := 0 // farthest index reachable so far
	for i, x := range nums {
		if i > reach {
			return false
		}
		reach = max(reach, i+x)
	}
	return true
}

func main() {
	assertEq(canJump([]int{2, 3, 1, 1, 4}), true)
	assertEq(canJump([]int{3, 2, 1, 0, 4}), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
