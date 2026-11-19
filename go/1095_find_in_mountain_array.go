//go:build ignore

// 1095. Find in Mountain Array (Hard)
// https://leetcode.com/problems/find-in-mountain-array/
// Run: go run 1095_find_in_mountain_array.go lc.go
package main

import (
	"fmt"
	"os"
)

type MountainArray struct { // provided by LeetCode (this stand-in counts calls; the limit is 100)
	a     []int
	calls int
}

func (this *MountainArray) get(index int) int {
	this.calls++
	return this.a[index]
}

func (this *MountainArray) length() int {
	return len(this.a)
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	n := mountainArr.length()
	lo, hi := 0, n-1
	for lo < hi { // find the peak
		mid := lo + (hi-lo)/2
		if mountainArr.get(mid) < mountainArr.get(mid+1) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	peak := lo
	search := func(lo, hi int, ascending bool) int {
		for lo <= hi {
			mid := lo + (hi-lo)/2
			v := mountainArr.get(mid)
			if v == target {
				return mid
			}
			if (v < target) == ascending {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		return -1
	}
	if i := search(0, peak, true); i != -1 {
		return i
	}
	return search(peak+1, n-1, false)
}

func main() {
	rng := func(from, to, step int) []int { // like Python's range
		out := []int{}
		for i := from; i != to; i += step {
			out = append(out, i)
		}
		return out
	}
	m := &MountainArray{a: []int{1, 2, 3, 4, 5, 3, 1}}
	assertEq(findInMountainArray(3, m), 2)
	assertEq(m.calls <= 100, true)
	assertEq(findInMountainArray(3, &MountainArray{a: []int{0, 1, 2, 4, 2, 1}}), -1)
	assertEq(findInMountainArray(1, &MountainArray{a: []int{0, 1, 2, 4, 2, 1}}), 1)
	big := &MountainArray{a: append(rng(0, 5000, 1), rng(4998, -1, -1)...)}
	assertEq(findInMountainArray(4000, big), 4000)
	assertEq(big.calls <= 100, true)
	assertEq(findInMountainArray(2001, &MountainArray{a: append(rng(0, 5000, 1), rng(4998, 2000, -1)...)}), 2001)
	assertEq(findInMountainArray(2000, &MountainArray{a: append(rng(10, 3000, 1), 5)}), 1990)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
