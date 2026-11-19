//go:build ignore

// 841. Keys and Rooms (Medium)
// https://leetcode.com/problems/keys-and-rooms/
// Run: go run 0841_keys_and_rooms.go lc.go
package main

import (
	"fmt"
	"os"
)

func canVisitAllRooms(rooms [][]int) bool {
	seen, st := map[int]bool{0: true}, []int{0}
	for len(st) > 0 {
		room := st[len(st)-1]
		st = st[:len(st)-1]
		for _, k := range rooms[room] {
			if !seen[k] {
				seen[k] = true
				st = append(st, k)
			}
		}
	}
	return len(seen) == len(rooms)
}

func main() {
	assertEq(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}), true)
	assertEq(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
