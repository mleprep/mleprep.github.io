//go:build ignore

// 721. Accounts Merge (Medium)
// https://leetcode.com/problems/accounts-merge/
// Run: go run 0721_accounts_merge.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	parent, owner := map[string]string{}, map[string]string{}
	var find func(x string) string
	find = func(x string) string {
		if _, ok := parent[x]; !ok {
			parent[x] = x
		}
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	for _, acc := range accounts { // union every email of an account with its first
		for _, e := range acc[1:] {
			owner[e] = acc[0]
			parent[find(e)] = find(acc[1])
		}
	}
	groups := map[string][]string{}
	for e := range owner {
		root := find(e)
		groups[root] = append(groups[root], e)
	}
	res := [][]string{}
	for root, emails := range groups {
		sort.Strings(emails)
		res = append(res, append([]string{owner[root]}, emails...))
	}
	return res
}

func main() {
	acc := [][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"},
		{"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}}
	assertUnordered(accountsMerge(acc), [][]string{
		{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
		{"John", "johnnybravo@mail.com"}, {"Mary", "mary@mail.com"}})
	assertUnordered(accountsMerge([][]string{{"Alex", "a@x.com"}, {"Alex", "b@x.com"}}), [][]string{{"Alex", "a@x.com"}, {"Alex", "b@x.com"}})
	assertEq(accountsMerge([][]string{{"Bo", "b@x.com", "a@x.com", "b@x.com"}}), [][]string{{"Bo", "a@x.com", "b@x.com"}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
