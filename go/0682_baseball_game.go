//go:build ignore

// 682. Baseball Game (Easy)
// https://leetcode.com/problems/baseball-game/
// Run: go run 0682_baseball_game.go lc.go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func calPoints(operations []string) int {
	st := []int{}
	for _, op := range operations {
		n := len(st)
		switch op {
		case "C":
			st = st[:n-1]
		case "D":
			st = append(st, 2*st[n-1])
		case "+":
			st = append(st, st[n-1]+st[n-2])
		default:
			v, _ := strconv.Atoi(op)
			st = append(st, v)
		}
	}
	sum := 0
	for _, x := range st {
		sum += x
	}
	return sum
}

func main() {
	assertEq(calPoints([]string{"5", "2", "C", "D", "+"}), 30)
	assertEq(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}), 27)
	assertEq(calPoints([]string{"1", "C"}), 0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
