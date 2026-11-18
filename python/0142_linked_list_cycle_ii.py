# 142. Linked List Cycle II (Medium)
# https://leetcode.com/problems/linked-list-cycle-ii/
# Run: python 0142_linked_list_cycle_ii.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        slow = fast = head
        while fast and fast.next:
            slow, fast = slow.next, fast.next.next
            if slow is fast:
                p = head
                while p is not slow:
                    p, slow = p.next, slow.next
                return p
        return None


if __name__ == "__main__":
    h = make_list([3, 2, 0, -4]); h.next.next.next.next = h.next
    assert Solution().detectCycle(h) is h.next
    assert Solution().detectCycle(make_list([1])) is None
    print("OK")
