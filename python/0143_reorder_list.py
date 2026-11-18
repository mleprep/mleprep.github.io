# 143. Reorder List (Medium)
# https://leetcode.com/problems/reorder-list/
# Run: python 0143_reorder_list.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        slow = fast = head
        while fast.next and fast.next.next:
            slow, fast = slow.next, fast.next.next
        prev, cur = None, slow.next
        slow.next = None
        while cur:
            cur.next, prev, cur = prev, cur, cur.next
        a, b = head, prev
        while b:
            an, bn = a.next, b.next
            a.next, b.next = b, an
            a, b = an, bn


if __name__ == "__main__":
    h = make_list([1, 2, 3, 4, 5]); Solution().reorderList(h)
    assert list_vals(h) == [1, 5, 2, 4, 3]
    h = make_list([1, 2, 3, 4]); Solution().reorderList(h)
    assert list_vals(h) == [1, 4, 2, 3]
    print("OK")
