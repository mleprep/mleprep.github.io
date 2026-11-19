//go:build ignore

// 355. Design Twitter (Medium)
// https://leetcode.com/problems/design-twitter/
// Run: go run 0355_design_twitter.go lc.go
package main

import (
	"container/heap"
	"fmt"
	"os"
)

type tweet struct{ time, id int }

type feedItem struct{ time, id, user, i int }
type feedHeap []feedItem // newest first

func (h feedHeap) Len() int           { return len(h) }
func (h feedHeap) Less(i, j int) bool { return h[i].time > h[j].time }
func (h feedHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *feedHeap) Push(x any)        { *h = append(*h, x.(feedItem)) }
func (h *feedHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type Twitter struct {
	t       int
	tweets  map[int][]tweet
	follows map[int]map[int]bool
}

func Constructor() Twitter {
	return Twitter{tweets: map[int][]tweet{}, follows: map[int]map[int]bool{}}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.t++
	this.tweets[userId] = append(this.tweets[userId], tweet{this.t, tweetId})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	h := &feedHeap{}
	users := []int{userId}
	for u := range this.follows[userId] {
		if u != userId {
			users = append(users, u)
		}
	}
	for _, u := range users { // merge each user's newest tweets, k-way
		if ts := this.tweets[u]; len(ts) > 0 {
			i := len(ts) - 1
			heap.Push(h, feedItem{ts[i].time, ts[i].id, u, i})
		}
	}
	res := []int{}
	for h.Len() > 0 && len(res) < 10 {
		it := heap.Pop(h).(feedItem)
		res = append(res, it.id)
		if it.i > 0 {
			prev := this.tweets[it.user][it.i-1]
			heap.Push(h, feedItem{prev.time, prev.id, it.user, it.i - 1})
		}
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.follows[followerId] == nil {
		this.follows[followerId] = map[int]bool{}
	}
	this.follows[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.follows[followerId], followeeId)
}

func main() {
	tw := Constructor()
	tw.PostTweet(1, 5)
	assertEq(tw.GetNewsFeed(1), []int{5})
	tw.Follow(1, 2)
	tw.PostTweet(2, 6)
	assertEq(tw.GetNewsFeed(1), []int{6, 5})
	tw.Unfollow(1, 2)
	assertEq(tw.GetNewsFeed(1), []int{5})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
