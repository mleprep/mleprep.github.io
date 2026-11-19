//go:build ignore

// 1462. Course Schedule IV (Medium)
// https://leetcode.com/problems/course-schedule-iv/
// Run: go run 1462_course_schedule_iv.go lc.go
package main

import (
	"fmt"
	"os"
)

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	n := numCourses
	reach := make([][]bool, n)
	for i := range reach {
		reach[i] = make([]bool, n)
	}
	for _, p := range prerequisites {
		reach[p[0]][p[1]] = true
	}
	for k := 0; k < n; k++ { // Floyd-Warshall transitive closure
		for i := 0; i < n; i++ {
			if reach[i][k] {
				for j := 0; j < n; j++ {
					if reach[k][j] {
						reach[i][j] = true
					}
				}
			}
		}
	}
	res := make([]bool, len(queries))
	for i, q := range queries {
		res[i] = reach[q[0]][q[1]]
	}
	return res
}

func main() {
	assertEq(checkIfPrerequisite(2, [][]int{{1, 0}}, [][]int{{0, 1}, {1, 0}}), []bool{false, true})
	assertEq(checkIfPrerequisite(2, [][]int{}, [][]int{{1, 0}, {0, 1}}), []bool{false, false})
	assertEq(checkIfPrerequisite(3, [][]int{{1, 2}, {1, 0}, {2, 0}}, [][]int{{1, 0}, {1, 2}}), []bool{true, true})
	assertEq(checkIfPrerequisite(4, [][]int{{0, 1}, {1, 2}, {2, 3}}, [][]int{{0, 3}, {3, 0}, {1, 3}}), []bool{true, false, true})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
