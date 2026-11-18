# 312. Burst Balloons (Hard)
# https://leetcode.com/problems/burst-balloons/
# Run: python 0312_burst_balloons.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maxCoins(self, nums: List[int]) -> int:
        a = [1] + nums + [1]
        n = len(a)
        dp = [[0] * n for _ in range(n)]
        for length in range(2, n):
            for l in range(n - length):
                r = l + length
                dp[l][r] = max(dp[l][k] + a[l] * a[k] * a[r] + dp[k][r] for k in range(l + 1, r))
        return dp[0][n - 1]


if __name__ == "__main__":
    assert Solution().maxCoins([3, 1, 5, 8]) == 167
    assert Solution().maxCoins([1, 5]) == 10
    print("OK")
