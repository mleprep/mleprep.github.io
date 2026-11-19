//go:build ignore

// 210. Course Schedule II (Medium)
// https://leetcode.com/problems/course-schedule-ii/
// Run: go run 0210_course_schedule_ii.go lc.go
package main

import (
	"fmt"
	"os"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	g, indeg := make([][]int, numCourses), make([]int, numCourses)
	for _, p := range prerequisites {
		g[p[1]] = append(g[p[1]], p[0])
		indeg[p[0]]++
	}
	q := []int{}
	for i, d := range indeg {
		if d == 0 {
			q = append(q, i)
		}
	}
	order := []int{}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		order = append(order, u)
		for _, v := range g[u] {
			if indeg[v]--; indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if len(order) < numCourses {
		return []int{}
	}
	return order
}

func main() {
	// any valid topological order is accepted
	valid := func(n int, pre [][]int, order []int) bool {
		pos := map[int]int{}
		for i, c := range order {
			pos[c] = i
		}
		if len(order) != n || len(pos) != n {
			return false
		}
		for _, p := range pre {
			if pos[p[1]] >= pos[p[0]] {
				return false
			}
		}
		return true
	}
	pre := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	assertEq(valid(4, pre, findOrder(4, pre)), true)
	assertEq(valid(2, [][]int{{1, 0}}, findOrder(2, [][]int{{1, 0}})), true)
	assertEq(findOrder(2, [][]int{{0, 1}, {1, 0}}), []int{})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
