//go:build ignore

// 735. Asteroid Collision (Medium)
// https://leetcode.com/problems/asteroid-collision/
// Run: go run 0735_asteroid_collision.go lc.go
package main

import (
	"fmt"
	"os"
)

func asteroidCollision(asteroids []int) []int {
	st := []int{}
	for _, x := range asteroids {
		alive := true
		for alive && x < 0 && len(st) > 0 && st[len(st)-1] > 0 {
			top := st[len(st)-1]
			if top < -x { // the stack top explodes, keep going
				st = st[:len(st)-1]
				continue
			}
			if top == -x {
				st = st[:len(st)-1]
			}
			alive = false
		}
		if alive {
			st = append(st, x)
		}
	}
	return st
}

func main() {
	assertEq(asteroidCollision([]int{5, 10, -5}), []int{5, 10})
	assertEq(asteroidCollision([]int{8, -8}), []int{})
	assertEq(asteroidCollision([]int{10, 2, -5}), []int{10})
	assertEq(asteroidCollision([]int{-2, -1, 1, 2}), []int{-2, -1, 1, 2})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
