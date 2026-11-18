# 141. Linked List Cycle (Easy)
# https://leetcode.com/problems/linked-list-cycle/
# Run: python 0141_linked_list_cycle.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        slow = fast = head
        while fast and fast.next:
            slow, fast = slow.next, fast.next.next
            if slow is fast:
                return True
        return False


if __name__ == "__main__":
    h = make_list([3, 2, 0, -4]); h.next.next.next.next = h.next   # tail -> node 1
    assert Solution().hasCycle(h) is True
    assert Solution().hasCycle(make_list([1])) is False
    print("OK")
