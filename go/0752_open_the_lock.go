//go:build ignore

// 752. Open the Lock (Medium)
// https://leetcode.com/problems/open-the-lock/
// Run: go run 0752_open_the_lock.go lc.go
package main

import (
	"fmt"
	"os"
)

func openLock(deadends []string, target string) int {
	dead := map[string]bool{}
	for _, d := range deadends {
		dead[d] = true
	}
	if dead["0000"] {
		return -1
	}
	q, seen := []string{"0000"}, map[string]bool{"0000": true}
	for steps := 0; len(q) > 0; steps++ {
		next := []string{}
		for _, cur := range q {
			if cur == target {
				return steps
			}
			b := []byte(cur)
			for i := 0; i < 4; i++ {
				orig := b[i]
				for _, delta := range []byte{1, 9} { // +1 or -1 (mod 10)
					b[i] = '0' + (orig-'0'+delta)%10
					if s := string(b); !seen[s] && !dead[s] {
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
	assertEq(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"), 6)
	assertEq(openLock([]string{"8888"}, "0009"), 1)
	assertEq(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"), -1)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
