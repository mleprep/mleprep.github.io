//go:build ignore

// 605. Can Place Flowers (Easy)
// https://leetcode.com/problems/can-place-flowers/
// Run: go run 0605_can_place_flowers.go lc.go
package main

import (
	"fmt"
	"os"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	bed := append(append([]int{0}, flowerbed...), 0) // pad both ends with empty plots
	for i := 1; i < len(bed)-1; i++ {
		if bed[i-1] == 0 && bed[i] == 0 && bed[i+1] == 0 { // planting as early as possible is optimal
			bed[i] = 1
			n--
		}
	}
	return n <= 0
}

func main() {
	assertEq(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1), true)
	assertEq(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2), false)
	assertEq(canPlaceFlowers([]int{0}, 1), true)
	assertEq(canPlaceFlowers([]int{0, 0, 1, 0, 0}, 2), true)
	assertEq(canPlaceFlowers([]int{1, 0, 1, 0, 1}, 0), true)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
