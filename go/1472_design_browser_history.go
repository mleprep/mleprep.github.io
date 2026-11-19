//go:build ignore

// 1472. Design Browser History (Medium)
// https://leetcode.com/problems/design-browser-history/
// Run: go run 1472_design_browser_history.go lc.go
package main

import (
	"fmt"
	"os"
)

type BrowserHistory struct {
	pages     []string
	cur, last int // current page; last valid page (later ones are stale)
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{pages: []string{homepage}}
}

func (this *BrowserHistory) Visit(url string) {
	this.cur++
	if this.cur == len(this.pages) {
		this.pages = append(this.pages, url)
	} else {
		this.pages[this.cur] = url // overwrite instead of truncating
	}
	this.last = this.cur // forward history is gone
}

func (this *BrowserHistory) Back(steps int) string {
	this.cur = max(0, this.cur-steps)
	return this.pages[this.cur]
}

func (this *BrowserHistory) Forward(steps int) string {
	this.cur = min(this.last, this.cur+steps)
	return this.pages[this.cur]
}

func main() {
	b := Constructor("leetcode.com")
	for _, u := range []string{"google.com", "facebook.com", "youtube.com"} {
		b.Visit(u)
	}
	assertEq([]string{b.Back(1), b.Back(1), b.Forward(1)}, []string{"facebook.com", "google.com", "facebook.com"})
	b.Visit("linkedin.com")
	assertEq([]string{b.Forward(2), b.Back(2), b.Back(7)}, []string{"linkedin.com", "google.com", "leetcode.com"})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
