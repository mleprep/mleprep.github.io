# 50. Pow(x, n) (Medium)
# https://leetcode.com/problems/powx-n/
# Run: python 0050_powx_n.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def myPow(self, x: float, n: int) -> float:
        if n < 0:
            x, n = 1 / x, -n
        res = 1.0
        while n:
            if n & 1:
                res *= x
            x *= x
            n >>= 1
        return res


if __name__ == "__main__":
    assert abs(Solution().myPow(2.0, 10) - 1024.0) < 1e-9
    assert abs(Solution().myPow(2.1, 3) - 9.261) < 1e-9
    assert abs(Solution().myPow(2.0, -2) - 0.25) < 1e-9
    print("OK")
