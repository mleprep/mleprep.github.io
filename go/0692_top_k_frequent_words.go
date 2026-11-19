//go:build ignore

// 692. Top K Frequent Words (Medium)
// https://leetcode.com/problems/top-k-frequent-words/
// Run: go run 0692_top_k_frequent_words.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func topKFrequent(words []string, k int) []string {
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w]++
	}
	keys := make([]string, 0, len(cnt))
	for w := range cnt {
		keys = append(keys, w)
	}
	sort.Slice(keys, func(i, j int) bool { // by count desc, then alphabetically
		a, b := keys[i], keys[j]
		if cnt[a] != cnt[b] {
			return cnt[a] > cnt[b]
		}
		return a < b
	})
	return keys[:k]
}

func main() {
	assertEq(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2), []string{"i", "love"})
	assertEq(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4), []string{"the", "is", "sunny", "day"})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
