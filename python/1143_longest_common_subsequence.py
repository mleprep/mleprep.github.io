# 1143. Longest Common Subsequence (Medium)
# https://leetcode.com/problems/longest-common-subsequence/
# Run: python 1143_longest_common_subsequence.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        n = len(text2)
        dp = [0] * (n + 1)
        for a in text1:
            diag = 0
            for j in range(1, n + 1):
                up = dp[j]
                dp[j] = diag + 1 if a == text2[j - 1] else max(dp[j], dp[j - 1])
                diag = up
        return dp[n]


if __name__ == "__main__":
    assert Solution().longestCommonSubsequence("abcde", "ace") == 3
    assert Solution().longestCommonSubsequence("abc", "def") == 0
    print("OK")
