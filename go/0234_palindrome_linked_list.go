//go:build ignore

// 234. Palindrome Linked List (Easy)
// https://leetcode.com/problems/palindrome-linked-list/
// Run: go run 0234_palindrome_linked_list.go lc.go
package main

import (
	"fmt"
	"os"
)

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	var prev *ListNode // reverse the second half, then compare
	for slow != nil {
		slow.Next, prev, slow = prev, slow, slow.Next
	}
	for ; prev != nil; prev, head = prev.Next, head.Next {
		if prev.Val != head.Val {
			return false
		}
	}
	return true
}

func main() {
	assertEq(isPalindrome(makeList(1, 2, 2, 1)), true)
	assertEq(isPalindrome(makeList(1, 2, 3, 2, 1)), true)
	assertEq(isPalindrome(makeList(1, 2)), false)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
