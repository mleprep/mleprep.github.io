# 516. Longest Palindromic Subsequence (Medium)
# https://leetcode.com/problems/longest-palindromic-subsequence/
# Run: python 0516_longest_palindromic_subsequence.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def longestPalindromeSubseq(self, s: str) -> int:
        n = len(s)
        dp = [0] * n                      # dp[r] for current l
        for l in range(n - 1, -1, -1):
            dp[l], diag = 1, 0            # diag = old dp[l+1][r-1]
            for r in range(l + 1, n):
                tmp = dp[r]
                dp[r] = diag + 2 if s[l] == s[r] else max(dp[r], dp[r - 1])
                diag = tmp
        return dp[n - 1]


if __name__ == "__main__":
    assert Solution().longestPalindromeSubseq("bbbab") == 4
    assert Solution().longestPalindromeSubseq("cbbd") == 2
    print("OK")
