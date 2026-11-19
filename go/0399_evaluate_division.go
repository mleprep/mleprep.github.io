//go:build ignore

// 399. Evaluate Division (Medium)
// https://leetcode.com/problems/evaluate-division/
// Run: go run 0399_evaluate_division.go lc.go
package main

import (
	"fmt"
	"os"
)

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	g := map[string]map[string]float64{} // a -> b -> a/b
	for i, e := range equations {
		a, b := e[0], e[1]
		if g[a] == nil {
			g[a] = map[string]float64{}
		}
		if g[b] == nil {
			g[b] = map[string]float64{}
		}
		g[a][b] = values[i]
		g[b][a] = 1 / values[i]
	}
	var dfs func(x, y string, seen map[string]bool) float64
	dfs = func(x, y string, seen map[string]bool) float64 { // product of weights along a path, or -1
		if x == y {
			return 1
		}
		seen[x] = true
		for nb, w := range g[x] {
			if !seen[nb] {
				if r := dfs(nb, y, seen); r != -1 {
					return w * r
				}
			}
		}
		return -1
	}
	res := make([]float64, len(queries))
	for i, q := range queries {
		if g[q[0]] == nil || g[q[1]] == nil {
			res[i] = -1
		} else {
			res[i] = dfs(q[0], q[1], map[string]bool{})
		}
	}
	return res
}

func main() {
	eq, val := [][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}
	qs := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	assertEq(calcEquation(eq, val, qs), []float64{6.0, 0.5, -1.0, 1.0, -1.0})
	assertEq(calcEquation([][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, []float64{1.5, 2.5, 5.0}, [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}), []float64{3.75, 0.4, 5.0, 0.2})
	assertEq(calcEquation([][]string{{"a", "b"}}, []float64{0.5}, [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}}), []float64{0.5, 2.0, -1.0, -1.0})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
