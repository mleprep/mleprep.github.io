# 234. Palindrome Linked List (Easy)
# https://leetcode.com/problems/palindrome-linked-list/
# Run: python 0234_palindrome_linked_list.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        slow = fast = head
        while fast and fast.next:
            slow, fast = slow.next, fast.next.next
        prev = None
        while slow:
            slow.next, prev, slow = prev, slow, slow.next
        while prev:
            if prev.val != head.val:
                return False
            prev, head = prev.next, head.next
        return True


if __name__ == "__main__":
    assert Solution().isPalindrome(make_list([1, 2, 2, 1])) is True
    assert Solution().isPalindrome(make_list([1, 2, 3, 2, 1])) is True
    assert Solution().isPalindrome(make_list([1, 2])) is False
    print("OK")
