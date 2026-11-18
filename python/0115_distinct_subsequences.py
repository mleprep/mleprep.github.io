# 115. Distinct Subsequences (Hard)
# https://leetcode.com/problems/distinct-subsequences/
# Run: python 0115_distinct_subsequences.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def numDistinct(self, s: str, t: str) -> int:
        dp = [1] + [0] * len(t)
        for ch in s:
            for j in range(len(t), 0, -1):
                if t[j - 1] == ch:
                    dp[j] += dp[j - 1]
        return dp[-1]


if __name__ == "__main__":
    assert Solution().numDistinct("rabbbit", "rabbit") == 3
    assert Solution().numDistinct("babgbag", "bag") == 5
    print("OK")
