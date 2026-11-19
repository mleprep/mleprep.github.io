# 86. Partition List (Medium)
# https://leetcode.com/problems/partition-list/
# Run: python 0086_partition_list.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        small = s = ListNode()
        big = b = ListNode()
        while head:
            if head.val < x:
                s.next = s = head
            else:
                b.next = b = head
            head = head.next
        b.next = None
        s.next = big.next
        return small.next


if __name__ == "__main__":
    assert list_vals(Solution().partition(make_list([1, 4, 3, 2, 5, 2]), 3)) == [1, 2, 2, 4, 3, 5]
    assert list_vals(Solution().partition(make_list([2, 1]), 2)) == [1, 2]
    assert list_vals(Solution().partition(make_list([]), 0)) == []
    print("OK")
