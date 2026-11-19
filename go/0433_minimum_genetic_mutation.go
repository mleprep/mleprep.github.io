//go:build ignore

// 433. Minimum Genetic Mutation (Medium)
// https://leetcode.com/problems/minimum-genetic-mutation/
// Run: go run 0433_minimum_genetic_mutation.go lc.go
package main

import (
	"fmt"
	"os"
)

func minMutation(startGene string, endGene string, bank []string) int {
	valid := map[string]bool{}
	for _, g := range bank {
		valid[g] = true
	}
	q, seen := []string{startGene}, map[string]bool{startGene: true}
	for steps := 0; len(q) > 0; steps++ {
		next := []string{}
		for _, g := range q {
			if g == endGene {
				return steps
			}
			b := []byte(g)
			for i := range b {
				orig := b[i]
				for _, ch := range []byte("ACGT") {
					b[i] = ch
					if s := string(b); valid[s] && !seen[s] {
						seen[s] = true
						next = append(next, s)
					}
				}
				b[i] = orig
			}
		}
		q = next
	}
	return -1
}

func main() {
	assertEq(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}), 1)
	assertEq(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}), 2)
	assertEq(minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}), 3)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
