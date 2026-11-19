# 474. Ones and Zeroes (Medium)
# https://leetcode.com/problems/ones-and-zeroes/
# Run: python 0474_ones_and_zeroes.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def findMaxForm(self, strs: List[str], m: int, n: int) -> int:
        dp = [[0] * (n + 1) for _ in range(m + 1)]
        for s in strs:
            z, o = s.count("0"), s.count("1")
            for i in range(m, z - 1, -1):
                for j in range(n, o - 1, -1):
                    dp[i][j] = max(dp[i][j], dp[i - z][j - o] + 1)
        return dp[m][n]


if __name__ == "__main__":
    assert Solution().findMaxForm(["10", "0001", "111001", "1", "0"], 5, 3) == 4
    assert Solution().findMaxForm(["10", "0", "1"], 1, 1) == 2
    print("OK")
