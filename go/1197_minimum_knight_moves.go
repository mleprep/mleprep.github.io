//go:build ignore

// 1197. Minimum Knight Moves (Medium)
// https://leetcode.com/problems/minimum-knight-moves/
// Run: go run 1197_minimum_knight_moves.go lc.go
package main

import (
	"fmt"
	"os"
)

func minKnightMoves(x int, y int) int {
	if x < 0 { // the board is symmetric
		x = -x
	}
	if y < 0 {
		y = -y
	}
	type state struct{ a, b int }
	seen := map[state]bool{{0, 0}: true}
	q := []state{{0, 0}}
	for d := 0; len(q) > 0; d++ {
		next := []state{}
		for _, s := range q {
			if s.a == x && s.b == y {
				return d
			}
			for _, m := range [][2]int{{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}} {
				n := state{s.a + m[0], s.b + m[1]}
				if n.a >= -2 && n.a <= x+2 && n.b >= -2 && n.b <= y+2 && !seen[n] { // a small margin is enough
					seen[n] = true
					next = append(next, n)
				}
			}
		}
		q = next
	}
	return -1
}

func main() {
	assertEq(minKnightMoves(2, 1), 1)
	assertEq(minKnightMoves(5, 5), 4)
	assertEq(minKnightMoves(0, 0), 0)
	assertEq(minKnightMoves(1, 1), 2)
	assertEq(minKnightMoves(1, 0), 3)
	assertEq(minKnightMoves(-1, -2), 1)
	assertEq(minKnightMoves(300, 0), 150)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
