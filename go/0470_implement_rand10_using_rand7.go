//go:build ignore

// 470. Implement Rand10() Using Rand7() (Medium)
// https://leetcode.com/problems/implement-rand10-using-rand7/
// Run: go run 0470_implement_rand10_using_rand7.go lc.go
package main

import (
	"fmt"
	"math/rand"
	"os"
)

var rng = rand.New(rand.NewSource(7)) // provided by LeetCode (a seeded stand-in)

func rand7() int { // provided by LeetCode
	return rng.Intn(7) + 1
}

func rand10() int {
	for {
		v := (rand7()-1)*7 + rand7() // uniform on 1..49
		if v <= 40 {                 // rejection sampling keeps 1..40, four of each final value
			return (v-1)%10 + 1
		}
	}
}

func main() {
	cnt := map[int]int{}
	for i := 0; i < 20000; i++ {
		cnt[rand10()]++
	}
	ok := len(cnt) == 10
	for v := 1; v <= 10; v++ {
		ok = ok && cnt[v] > 1800 && cnt[v] < 2200
	}
	assertEq(ok, true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
