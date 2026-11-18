# 70. Climbing Stairs (Easy)
# https://leetcode.com/problems/climbing-stairs/
# Run: python 0070_climbing_stairs.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def climbStairs(self, n: int) -> int:
        a, b = 1, 1          # ways(0), ways(1)
        for _ in range(n - 1):
            a, b = b, a + b
        return b


if __name__ == "__main__":
    assert Solution().climbStairs(2) == 2
    assert Solution().climbStairs(5) == 8
    print("OK")
