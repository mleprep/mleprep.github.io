//go:build ignore

// 227. Basic Calculator II (Medium)
// https://leetcode.com/problems/basic-calculator-ii/
// Run: go run 0227_basic_calculator_ii.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func calculate(s string) int {
	st, num, op := []int{}, 0, byte('+')
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
		}
		if strings.IndexByte("+-*/", ch) >= 0 || i == len(s)-1 {
			switch op { // apply the previous operator to num
			case '+':
				st = append(st, num)
			case '-':
				st = append(st, -num)
			case '*':
				st[len(st)-1] *= num
			case '/':
				st[len(st)-1] /= num
			}
			op, num = ch, 0
		}
	}
	sum := 0
	for _, x := range st {
		sum += x
	}
	return sum
}

func main() {
	assertEq(calculate("3+2*2"), 7)
	assertEq(calculate(" 3/2 "), 1)
	assertEq(calculate(" 3+5 / 2 "), 5)
	assertEq(calculate("14-3/2"), 13)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
