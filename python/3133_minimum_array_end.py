# 3133. Minimum Array End (Medium)
# https://leetcode.com/problems/minimum-array-end/
# Run: python 3133_minimum_array_end.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minEnd(self, n: int, x: int) -> int:
        res, k, bit = x, n - 1, 1
        while k:
            if not x & bit:              # free position: place the next bit of n - 1
                if k & 1:
                    res |= bit
                k >>= 1
            bit <<= 1
        return res


if __name__ == "__main__":
    assert Solution().minEnd(3, 4) == 6
    assert Solution().minEnd(2, 7) == 15
    assert Solution().minEnd(1, 5) == 5
    assert Solution().minEnd(5, 1) == 9
    def brute(n, x):                     # n-th smallest number containing all bits of x
        v, seen = x, 0
        while True:
            if v & x == x:
                seen += 1
                if seen == n:
                    return v
            v += 1
    assert all(Solution().minEnd(n, x) == brute(n, x) for n in range(1, 25) for x in range(1, 33))
    print("OK")
