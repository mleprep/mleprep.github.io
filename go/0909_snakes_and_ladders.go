//go:build ignore

// 909. Snakes and Ladders (Medium)
// https://leetcode.com/problems/snakes-and-ladders/
// Run: go run 0909_snakes_and_ladders.go lc.go
package main

import (
	"fmt"
	"os"
)

func snakesAndLadders(board [][]int) int {
	n := len(board)
	cell := func(s int) int { // board value of square s (boustrophedon from the bottom-left)
		r, c := (s-1)/n, (s-1)%n
		if r%2 == 1 {
			c = n - 1 - c
		}
		return board[n-1-r][c]
	}
	q, seen := []int{1}, map[int]bool{1: true}
	for moves := 0; len(q) > 0; moves++ {
		next := []int{}
		for _, s := range q {
			if s == n*n {
				return moves
			}
			for nxt := s + 1; nxt <= min(s+6, n*n); nxt++ {
				dest := nxt
				if v := cell(nxt); v != -1 {
					dest = v
				}
				if !seen[dest] {
					seen[dest] = true
					next = append(next, dest)
				}
			}
		}
		q = next
	}
	return -1
}

func main() {
	B := [][]int{{-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 15, -1, -1, -1, -1}}
	assertEq(snakesAndLadders(B), 4)
	assertEq(snakesAndLadders([][]int{{-1, -1}, {-1, 3}}), 1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
