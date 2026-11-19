# 61. Rotate List (Medium)
# https://leetcode.com/problems/rotate-list/
# Run: python 0061_rotate_list.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        if not head:
            return None
        n, tail = 1, head
        while tail.next:
            tail, n = tail.next, n + 1
        tail.next = head                      # ring
        for _ in range(n - k % n - 1):
            head = head.next
        new_head, head.next = head.next, None
        return new_head


if __name__ == "__main__":
    assert list_vals(Solution().rotateRight(make_list([1, 2, 3, 4, 5]), 2)) == [4, 5, 1, 2, 3]
    assert list_vals(Solution().rotateRight(make_list([0, 1, 2]), 4)) == [2, 0, 1]
    print("OK")
