# 24. Swap Nodes in Pairs (Medium)
# https://leetcode.com/problems/swap-nodes-in-pairs/
# Run: python 0024_swap_nodes_in_pairs.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = prev = ListNode(0, head)
        while prev.next and prev.next.next:
            a, b = prev.next, prev.next.next
            prev.next, a.next, b.next = b, b.next, a
            prev = a
        return dummy.next


if __name__ == "__main__":
    assert list_vals(Solution().swapPairs(make_list([1, 2, 3, 4]))) == [2, 1, 4, 3]
    assert list_vals(Solution().swapPairs(make_list([1, 2, 3]))) == [2, 1, 3]
    print("OK")
