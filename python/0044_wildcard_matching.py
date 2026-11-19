# 44. Wildcard Matching (Hard)
# https://leetcode.com/problems/wildcard-matching/
# Run: python 0044_wildcard_matching.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        n = len(p)
        dp = [True] + [False] * n
        for j in range(1, n + 1):
            dp[j] = dp[j - 1] and p[j - 1] == "*"
        for ch in s:
            diag, dp[0] = dp[0], False
            for j in range(1, n + 1):
                up = dp[j]
                if p[j - 1] == "*":
                    dp[j] = dp[j - 1] or up
                else:
                    dp[j] = diag and p[j - 1] in (ch, "?")
                diag = up
        return dp[n]


if __name__ == "__main__":
    assert Solution().isMatch("aa", "a") is False
    assert Solution().isMatch("aa", "*") is True
    assert Solution().isMatch("cb", "?a") is False
    assert Solution().isMatch("adceb", "*a*b") is True
    assert Solution().isMatch("acdcb", "a*c?b") is False
    print("OK")
