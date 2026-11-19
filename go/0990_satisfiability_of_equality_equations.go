//go:build ignore

// 990. Satisfiability of Equality Equations (Medium)
// https://leetcode.com/problems/satisfiability-of-equality-equations/
// Run: go run 0990_satisfiability_of_equality_equations.go lc.go
package main

import (
	"fmt"
	"os"
)

func equationsPossible(equations []string) bool {
	var parent [26]int
	for i := range parent {
		parent[i] = i
	}
	find := func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	for _, e := range equations { // union all "==" first
		if e[1] == '=' {
			parent[find(int(e[0]-'a'))] = find(int(e[3] - 'a'))
		}
	}
	for _, e := range equations { // then every "!=" must span two groups
		if e[1] == '!' && find(int(e[0]-'a')) == find(int(e[3]-'a')) {
			return false
		}
	}
	return true
}

func main() {
	assertEq(equationsPossible([]string{"a==b", "b!=a"}), false)
	assertEq(equationsPossible([]string{"b==a", "a==b"}), true)
	assertEq(equationsPossible([]string{"a==b", "b==c", "a==c"}), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
