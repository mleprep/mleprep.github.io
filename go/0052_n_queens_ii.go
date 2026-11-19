//go:build ignore

// 52. N-Queens II (Hard)
// https://leetcode.com/problems/n-queens-ii/
// Run: go run 0052_n_queens_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func totalNQueens(n int) int {
	full := 1<<n - 1
	var bt func(cols, d1, d2 int) int // bitmasks of attacked columns in the current row
	bt = func(cols, d1, d2 int) int {
		if cols == full {
			return 1
		}
		count := 0
		for free := ^(cols | d1 | d2) & full; free != 0; {
			bit := free & -free // lowest free column
			free -= bit
			count += bt(cols|bit, (d1|bit)<<1&full, (d2|bit)>>1)
		}
		return count
	}
	return bt(0, 0, 0)
}

func main() {
	assertEq(totalNQueens(4), 2)
	assertEq(totalNQueens(1), 1)
	assertEq(totalNQueens(8), 92)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
