//go:build ignore

// 282. Expression Add Operators (Hard)
// https://leetcode.com/problems/expression-add-operators/
// Run: go run 0282_expression_add_operators.go lc.go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func addOperators(num string, target int) []string {
	res := []string{}
	// prev is the last multiplicative term, so "*" can undo it: value - prev + prev*cur
	var bt func(i int, expr string, value, prev int)
	bt = func(i int, expr string, value, prev int) {
		if i == len(num) {
			if value == target {
				res = append(res, expr)
			}
			return
		}
		for j := i + 1; j <= len(num); j++ {
			s := num[i:j]
			if len(s) > 1 && s[0] == '0' { // no leading zeros
				break
			}
			cur, _ := strconv.Atoi(s)
			if i == 0 {
				bt(j, s, cur, cur)
			} else {
				bt(j, expr+"+"+s, value+cur, cur)
				bt(j, expr+"-"+s, value-cur, -cur)
				bt(j, expr+"*"+s, value-prev+prev*cur, prev*cur)
			}
		}
	}
	bt(0, "", 0, 0)
	return res
}

func main() {
	assertUnordered(addOperators("123", 6), []string{"1*2*3", "1+2+3"})
	assertUnordered(addOperators("232", 8), []string{"2*3+2", "2+3*2"})
	assertEq(addOperators("3456237490", 9191), []string{})
	assertUnordered(addOperators("105", 5), []string{"1*0+5", "10-5"})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
