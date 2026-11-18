# 72. Edit Distance (Medium)
# https://leetcode.com/problems/edit-distance/
# Run: python 0072_edit_distance.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        n = len(word2)
        dp = list(range(n + 1))
        for i, a in enumerate(word1, 1):
            diag, dp[0] = dp[0], i
            for j in range(1, n + 1):
                up = dp[j]
                dp[j] = diag if a == word2[j - 1] else 1 + min(diag, up, dp[j - 1])
                diag = up
        return dp[n]


if __name__ == "__main__":
    assert Solution().minDistance("horse", "ros") == 3
    assert Solution().minDistance("intention", "execution") == 5
    print("OK")
