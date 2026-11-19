# 343. Integer Break (Medium)
# https://leetcode.com/problems/integer-break/
# Run: python 0343_integer_break.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def integerBreak(self, n: int) -> int:
        dp = [0, 1] + [0] * (n - 1)
        for i in range(2, n + 1):
            dp[i] = max(max(j, dp[j]) * (i - j) for j in range(1, i))
        return dp[n]


if __name__ == "__main__":
    assert Solution().integerBreak(2) == 1
    assert Solution().integerBreak(10) == 36
    print("OK")
