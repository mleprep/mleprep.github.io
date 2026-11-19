//go:build ignore

// 93. Restore IP Addresses (Medium)
// https://leetcode.com/problems/restore-ip-addresses/
// Run: go run 0093_restore_ip_addresses.go lc.go
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res := []string{}
	var bt func(i int, parts []string)
	bt = func(i int, parts []string) {
		if len(parts) == 4 {
			if i == len(s) {
				res = append(res, strings.Join(parts, "."))
			}
			return
		}
		rem := 4 - len(parts)
		if left := len(s) - i; left < rem || left > 3*rem { // too few or too many digits left
			return
		}
		for L := 1; L <= 3 && i+L <= len(s); L++ {
			p := s[i : i+L]
			if v, _ := strconv.Atoi(p); (p[0] == '0' && L > 1) || v > 255 {
				break
			}
			bt(i+L, append(parts, p))
		}
	}
	bt(0, []string{})
	return res
}

func main() {
	assertUnordered(restoreIpAddresses("25525511135"), []string{"255.255.11.135", "255.255.111.35"})
	assertUnordered(restoreIpAddresses("0000"), []string{"0.0.0.0"})
	assertUnordered(restoreIpAddresses("101023"), []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
