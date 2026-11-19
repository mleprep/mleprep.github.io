//go:build ignore

// 901. Online Stock Span (Medium)
// https://leetcode.com/problems/online-stock-span/
// Run: go run 0901_online_stock_span.go lc.go
package main

import (
	"fmt"
	"os"
)

type StockSpanner struct {
	st [][2]int // (price, span)
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	span := 1
	for n := len(this.st); n > 0 && this.st[n-1][0] <= price; n = len(this.st) {
		span += this.st[n-1][1]
		this.st = this.st[:n-1]
	}
	this.st = append(this.st, [2]int{price, span})
	return span
}

func main() {
	s := Constructor()
	got := []int{}
	for _, p := range []int{100, 80, 60, 70, 60, 75, 85} {
		got = append(got, s.Next(p))
	}
	assertEq(got, []int{1, 1, 1, 2, 1, 4, 6})
	s2 := Constructor()
	got2 := []int{}
	for _, p := range []int{31, 41, 48, 59, 79} {
		got2 = append(got2, s2.Next(p))
	}
	assertEq(got2, []int{1, 2, 3, 4, 5})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
