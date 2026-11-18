# 206. Reverse Linked List (Easy)
# https://leetcode.com/problems/reverse-linked-list/
# Run: python 0206_reverse_linked_list.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        prev, cur = None, head
        while cur:
            cur.next, prev, cur = prev, cur, cur.next
        return prev


if __name__ == "__main__":
    assert list_vals(Solution().reverseList(make_list([1, 2, 3, 4, 5]))) == [5, 4, 3, 2, 1]
    assert Solution().reverseList(None) is None
    print("OK")
