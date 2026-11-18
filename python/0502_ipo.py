# 502. IPO (Hard)
# https://leetcode.com/problems/ipo/
# Run: python 0502_ipo.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findMaximizedCapital(self, k: int, w: int, profits: List[int], capital: List[int]) -> int:
        projects = sorted(zip(capital, profits))
        h, i = [], 0
        for _ in range(k):
            while i < len(projects) and projects[i][0] <= w:
                heappush(h, -projects[i][1])
                i += 1
            if not h:
                break
            w -= heappop(h)
        return w


if __name__ == "__main__":
    assert Solution().findMaximizedCapital(2, 0, [1, 2, 3], [0, 1, 1]) == 4
    assert Solution().findMaximizedCapital(3, 0, [1, 2, 3], [0, 1, 2]) == 6
    print("OK")
