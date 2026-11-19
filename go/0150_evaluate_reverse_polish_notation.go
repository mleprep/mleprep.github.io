//go:build ignore

// 150. Evaluate Reverse Polish Notation (Medium)
// https://leetcode.com/problems/evaluate-reverse-polish-notation/
// Run: go run 0150_evaluate_reverse_polish_notation.go lc.go
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evalRPN(tokens []string) int {
	st := []int{}
	for _, t := range tokens {
		if len(t) == 1 && strings.Contains("+-*/", t) {
			b, a := st[len(st)-1], st[len(st)-2]
			st = st[:len(st)-2]
			switch t {
			case "+":
				st = append(st, a+b)
			case "-":
				st = append(st, a-b)
			case "*":
				st = append(st, a*b)
			default:
				st = append(st, a/b) // Go's / truncates toward 0, as required
			}
		} else {
			n, _ := strconv.Atoi(t)
			st = append(st, n)
		}
	}
	return st[0]
}

func main() {
	assertEq(evalRPN([]string{"2", "1", "+", "3", "*"}), 9)
	assertEq(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}), 22)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
