//go:build ignore

// 118. Pascal's Triangle (Easy)
// https://leetcode.com/problems/pascals-triangle/
// Run: go run 0118_pascals_triangle.go lc.go
package main

import (
	"fmt"
	"os"
)

func generate(numRows int) [][]int {
	rows := [][]int{{1}}
	for len(rows) < numRows {
		prev := rows[len(rows)-1]
		row := []int{1}
		for i := 0; i+1 < len(prev); i++ {
			row = append(row, prev[i]+prev[i+1])
		}
		rows = append(rows, append(row, 1))
	}
	return rows
}

func main() {
	assertEq(generate(5), [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}})
	assertEq(generate(1), [][]int{{1}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
