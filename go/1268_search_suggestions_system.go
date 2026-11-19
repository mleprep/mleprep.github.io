//go:build ignore

// 1268. Search Suggestions System (Medium)
// https://leetcode.com/problems/search-suggestions-system/
// Run: go run 1268_search_suggestions_system.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	res := [][]string{}
	for k := 1; k <= len(searchWord); k++ {
		prefix := searchWord[:k]
		i := sort.SearchStrings(products, prefix) // first product >= prefix
		group := []string{}
		for j := i; j < len(products) && j < i+3 && strings.HasPrefix(products[j], prefix); j++ {
			group = append(group, products[j])
		}
		res = append(res, group)
	}
	return res
}

func main() {
	P := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	assertEq(suggestedProducts(P, "mouse"), [][]string{
		{"mobile", "moneypot", "monitor"}, {"mobile", "moneypot", "monitor"},
		{"mouse", "mousepad"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}})
	assertEq(suggestedProducts([]string{"havana"}, "tatiana"), [][]string{{}, {}, {}, {}, {}, {}, {}})
	assertEq(suggestedProducts([]string{"bags", "baggage", "banner", "box", "cloths"}, "bags"), [][]string{{"baggage", "bags", "banner"}, {"baggage", "bags", "banner"}, {"baggage", "bags"}, {"bags"}})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
