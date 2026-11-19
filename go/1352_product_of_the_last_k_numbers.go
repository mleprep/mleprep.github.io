//go:build ignore

// 1352. Product of the Last K Numbers (Medium)
// https://leetcode.com/problems/product-of-the-last-k-numbers/
// Run: go run 1352_product_of_the_last_k_numbers.go lc.go
package main

import (
	"fmt"
	"os"
)

type ProductOfNumbers struct {
	pre []int // prefix products since the last zero
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{pre: []int{1}}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.pre = []int{1} // any window containing the zero has product 0
	} else {
		this.pre = append(this.pre, this.pre[len(this.pre)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.pre)
	if k >= n {
		return 0 // the window reaches back past a zero
	}
	return this.pre[n-1] / this.pre[n-1-k]
}

func main() {
	p := Constructor()
	for _, x := range []int{3, 0, 2, 5, 4} {
		p.Add(x)
	}
	assertEq([]int{p.GetProduct(2), p.GetProduct(3), p.GetProduct(4)}, []int{20, 40, 0})
	p.Add(8)
	assertEq(p.GetProduct(2), 32)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
