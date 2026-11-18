# 221. Maximal Square (Medium)
# https://leetcode.com/problems/maximal-square/
# Run: python 0221_maximal_square.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def maximalSquare(self, matrix: List[List[str]]) -> int:
        n = len(matrix[0])
        dp, best = [0] * (n + 1), 0
        for row in matrix:
            diag = 0                          # dp[r-1][c-1]
            for c in range(1, n + 1):
                up = dp[c]
                dp[c] = 1 + min(dp[c], dp[c - 1], diag) if row[c - 1] == "1" else 0
                diag = up
                best = max(best, dp[c])
        return best * best


if __name__ == "__main__":
    M = ["10100", "10111", "11111", "10010"]
    assert Solution().maximalSquare([list(r) for r in M]) == 4
    assert Solution().maximalSquare([["0"]]) == 0
    print("OK")
