# 25. Reverse Nodes in k-Group (Hard)
# https://leetcode.com/problems/reverse-nodes-in-k-group/
# Run: python 0025_reverse_nodes_in_k_group.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        dummy = ListNode(0, head)
        group_prev = dummy
        while True:
            kth = group_prev
            for _ in range(k):
                kth = kth.next
                if not kth:
                    return dummy.next
            group_next = kth.next
            prev, cur = group_next, group_prev.next      # reverse [group_prev.next .. kth]
            while cur is not group_next:
                cur.next, prev, cur = prev, cur, cur.next
            first = group_prev.next
            group_prev.next = kth
            group_prev = first


if __name__ == "__main__":
    assert list_vals(Solution().reverseKGroup(make_list([1, 2, 3, 4, 5]), 2)) == [2, 1, 4, 3, 5]
    assert list_vals(Solution().reverseKGroup(make_list([1, 2, 3, 4, 5]), 3)) == [3, 2, 1, 4, 5]
    print("OK")
