//go:build ignore

// 207. Course Schedule (Medium)
// https://leetcode.com/problems/course-schedule/
// Run: go run 0207_course_schedule.go lc.go
package main

import (
	"fmt"
	"os"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
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
	taken := 0
	for len(q) > 0 { // Kahn's algorithm: a cycle leaves courses with indegree > 0
		u := q[0]
		q = q[1:]
		taken++
		for _, v := range g[u] {
			if indeg[v]--; indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return taken == numCourses
}

func main() {
	assertEq(canFinish(2, [][]int{{1, 0}}), true)
	assertEq(canFinish(2, [][]int{{1, 0}, {0, 1}}), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
