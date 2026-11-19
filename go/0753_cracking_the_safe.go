//go:build ignore

// 753. Cracking the Safe (Hard)
// https://leetcode.com/problems/cracking-the-safe/
// Run: go run 0753_cracking_the_safe.go lc.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func crackSafe(n int, k int) string {
	seen := map[string]bool{} // edges (n-digit codes) already used
	out := []byte{}
	var dfs func(node string)
	dfs = func(node string) { // Hierholzer on the de Bruijn graph of (n-1)-digit nodes
		for d := 0; d < k; d++ {
			edge := node + string(rune('0'+d))
			if !seen[edge] {
				seen[edge] = true
				dfs(edge[1:])
				out = append(out, byte('0'+d))
			}
		}
	}
	start := strings.Repeat("0", n-1)
	dfs(start)
	return string(out) + start
}

func main() {
	covers := func(s string, n, k int) bool { // every n-digit code appears, in the shortest length
		codes := map[string]bool{}
		for i := 0; i+n <= len(s); i++ {
			codes[s[i:i+n]] = true
		}
		want := 1
		for i := 0; i < n; i++ {
			want *= k
		}
		return len(codes) == want && len(s) == want+n-1
	}
	assertEq(covers(crackSafe(1, 2), 1, 2), true)
	assertEq(covers(crackSafe(2, 2), 2, 2), true)
	assertEq(covers(crackSafe(3, 3), 3, 3), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
