//go:build ignore

// 224. Basic Calculator (Hard)
// https://leetcode.com/problems/basic-calculator/
// Run: go run 0224_basic_calculator.go lc.go
package main

import (
	"fmt"
	"os"
)

func calculate(s string) int {
	res, num, sign := 0, 0, 1
	st := [][2]int{} // (result before "(", sign before "(")
	for i := 0; i < len(s); i++ {
		switch ch := s[i]; {
		case ch >= '0' && ch <= '9':
			num = num*10 + int(ch-'0')
		case ch == '+' || ch == '-':
			res += sign * num
			num, sign = 0, 1
			if ch == '-' {
				sign = -1
			}
		case ch == '(':
			st = append(st, [2]int{res, sign})
			res, sign = 0, 1
		case ch == ')':
			res += sign * num
			num = 0
			top := st[len(st)-1]
			st = st[:len(st)-1]
			res = top[0] + top[1]*res
		}
	}
	return res + sign*num
}

func main() {
	assertEq(calculate("1 + 1"), 2)
	assertEq(calculate(" 2-1 + 2 "), 3)
	assertEq(calculate("(1+(4+5+2)-3)+(6+8)"), 23)
	assertEq(calculate("-(2+3)"), -5)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
