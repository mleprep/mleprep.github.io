//go:build ignore

// 567. Permutation in String (Medium)
// https://leetcode.com/problems/permutation-in-string/
// Run: go run 0567_permutation_in_string.go lc.go
package main

import (
	"fmt"
	"os"
)

func checkInclusion(s1 string, s2 string) bool {
	k := len(s1)
	if k > len(s2) {
		return false
	}
	var need, win [26]int
	for i := 0; i < k; i++ {
		need[s1[i]-'a']++
		win[s2[i]-'a']++
	}
	if need == win {
		return true
	}
	for r := k; r < len(s2); r++ {
		win[s2[r]-'a']++
		win[s2[r-k]-'a']--
		if win == need {
			return true
		}
	}
	return false
}

func main() {
	assertEq(checkInclusion("ab", "eidbaooo"), true)
	assertEq(checkInclusion("ab", "eidboaoo"), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
