//go:build ignore

// 45. Jump Game II (Medium)
// https://leetcode.com/problems/jump-game-ii/
// Run: go run 0045_jump_game_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func jump(nums []int) int {
	jumps, end, far := 0, 0, 0 // BFS by levels: [.., end] is reachable with `jumps` jumps
	for i := 0; i < len(nums)-1; i++ {
		far = max(far, i+nums[i])
		if i == end {
			jumps++
			end = far
		}
	}
	return jumps
}

func main() {
	assertEq(jump([]int{2, 3, 1, 1, 4}), 2)
	assertEq(jump([]int{2, 3, 0, 1, 4}), 2)
	assertEq(jump([]int{0}), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
