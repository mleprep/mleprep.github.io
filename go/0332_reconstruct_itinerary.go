//go:build ignore

// 332. Reconstruct Itinerary (Hard)
// https://leetcode.com/problems/reconstruct-itinerary/
// Run: go run 0332_reconstruct_itinerary.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	g := map[string][]string{}
	for _, t := range tickets {
		g[t[0]] = append(g[t[0]], t[1])
	}
	for _, dests := range g {
		sort.Sort(sort.Reverse(sort.StringSlice(dests))) // popping from the end gives the smallest
	}
	route, st := []string{}, []string{"JFK"}
	for len(st) > 0 { // Hierholzer: an airport joins the route once it has no tickets left
		top := st[len(st)-1]
		if d := g[top]; len(d) > 0 {
			g[top] = d[:len(d)-1]
			st = append(st, d[len(d)-1])
		} else {
			route = append(route, top)
			st = st[:len(st)-1]
		}
	}
	for i, j := 0, len(route)-1; i < j; i, j = i+1, j-1 {
		route[i], route[j] = route[j], route[i]
	}
	return route
}

func main() {
	T := [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}
	assertEq(findItinerary(T), []string{"JFK", "MUC", "LHR", "SFO", "SJC"})
	assertEq(findItinerary([][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}), []string{"JFK", "NRT", "JFK", "KUL"})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
