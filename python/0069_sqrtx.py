# 69. Sqrt(x) (Easy)
# https://leetcode.com/problems/sqrtx/
# Run: python 0069_sqrtx.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def mySqrt(self, x: int) -> int:
        lo, hi = 0, x + 1          # first m with m*m > x
        while lo < hi:
            m = (lo + hi) // 2
            if m * m > x:
                hi = m
            else:
                lo = m + 1
        return lo - 1


if __name__ == "__main__":
    assert Solution().mySqrt(4) == 2
    assert Solution().mySqrt(8) == 2
    assert Solution().mySqrt(0) == 0
    print("OK")
