//go:build ignore

// 767. Reorganize String (Medium)
// https://leetcode.com/problems/reorganize-string/
// Run: go run 0767_reorganize_string.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"os"
)

type charCount struct {
	n  int
	ch byte
}
type countHeap []charCount // most frequent first

func (h countHeap) Len() int           { return len(h) }
func (h countHeap) Less(i, j int) bool { return h[i].n > h[j].n }
func (h countHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *countHeap) Push(x any)        { *h = append(*h, x.(charCount)) }
func (h *countHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func reorganizeString(s string) string {
	var cnt [26]int
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	h := &countHeap{}
	for i, n := range cnt {
		if n > 0 {
			heap.Push(h, charCount{n, byte('a' + i)})
		}
	}
	res := []byte{}
	var prev *charCount // the letter just used waits one turn
	for h.Len() > 0 {
		cur := heap.Pop(h).(charCount)
		res = append(res, cur.ch)
		if prev != nil {
			heap.Push(h, *prev)
			prev = nil
		}
		if cur.n--; cur.n > 0 {
			prev = &cur
		}
	}
	if prev != nil { // a letter is left over with nothing to separate it
		return ""
	}
	return string(res)
}

func main() {
	assertEq(reorganizeString("aab"), "aba")
	assertEq(reorganizeString("aaab"), "")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
