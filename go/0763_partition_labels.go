//go:build ignore

// 763. Partition Labels (Medium)
// https://leetcode.com/problems/partition-labels/
// Run: go run 0763_partition_labels.go lc.go
package main

import (
	"fmt"
	"os"
)

func partitionLabels(s string) []int {
	last := [26]int{}
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}
	res := []int{}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		end = max(end, last[s[i]-'a']) // the part must reach every letter's last occurrence
		if i == end {
			res = append(res, end-start+1)
			start = i + 1
		}
	}
	return res
}

func main() {
	assertEq(partitionLabels("ababcbacadefegdehijhklij"), []int{9, 7, 8})
	assertEq(partitionLabels("eccbbbbdec"), []int{10})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
