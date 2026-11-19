//go:build ignore

// 2392. Build a Matrix With Conditions (Hard)
// https://leetcode.com/problems/build-a-matrix-with-conditions/
// Run: go run 2392_build_a_matrix_with_conditions.go lc.go
package main

import (
	"fmt"
	"os"
)

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	topo := func(conds [][]int) []int { // pos[v] = v's place in a topological order; nil on a cycle
		adj := make([][]int, k+1)
		indeg := make([]int, k+1)
		for _, c := range conds {
			adj[c[0]] = append(adj[c[0]], c[1])
			indeg[c[1]]++
		}
		q := []int{}
		for v := 1; v <= k; v++ {
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
		pos := make([]int, k+1)
		placed := 0
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			pos[v] = placed
			placed++
			for _, w := range adj[v] {
				indeg[w]--
				if indeg[w] == 0 {
					q = append(q, w)
				}
			}
		}
		if placed < k {
			return nil
		}
		return pos
	}
	row, col := topo(rowConditions), topo(colConditions)
	if row == nil || col == nil {
		return [][]int{}
	}
	m := make([][]int, k)
	for i := range m {
		m[i] = make([]int, k)
	}
	for v := 1; v <= k; v++ {
		m[row[v]][col[v]] = v
	}
	return m
}

func main() {
	valid := func(k int, rc, cc, m [][]int) bool { // any matrix satisfying the conditions is accepted
		if len(m) != k {
			return false
		}
		row, col, cells := map[int]int{}, map[int]int{}, 0
		for i, r := range m {
			if len(r) != k {
				return false
			}
			for j, v := range r {
				if v != 0 {
					row[v], col[v] = i, j
					cells++
				}
			}
		}
		for v := 1; v <= k; v++ {
			if _, ok := row[v]; !ok {
				return false
			}
		}
		for _, c := range rc {
			if row[c[0]] >= row[c[1]] {
				return false
			}
		}
		for _, c := range cc {
			if col[c[0]] >= col[c[1]] {
				return false
			}
		}
		return cells == k
	}
	rc, cc := [][]int{{1, 2}, {3, 2}}, [][]int{{2, 1}, {3, 2}}
	assertEq(valid(3, rc, cc, buildMatrix(3, rc, cc)), true)
	assertEq(buildMatrix(3, [][]int{{1, 2}, {2, 3}, {3, 1}, {2, 3}}, [][]int{{2, 1}}), [][]int{})
	rc, cc = [][]int{{4, 3}, {3, 2}, {2, 1}}, [][]int{{1, 2}, {2, 3}}
	assertEq(valid(4, rc, cc, buildMatrix(4, rc, cc)), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
