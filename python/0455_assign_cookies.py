# 455. Assign Cookies (Easy)
# https://leetcode.com/problems/assign-cookies/
# Run: python 0455_assign_cookies.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findContentChildren(self, g: List[int], s: List[int]) -> int:
        g.sort(); s.sort()
        child = 0
        for size in s:
            if child < len(g) and size >= g[child]:
                child += 1
        return child


if __name__ == "__main__":
    assert Solution().findContentChildren([1, 2, 3], [1, 1]) == 1
    assert Solution().findContentChildren([1, 2], [1, 2, 3]) == 2
    print("OK")
