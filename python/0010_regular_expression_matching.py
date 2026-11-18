# 10. Regular Expression Matching (Hard)
# https://leetcode.com/problems/regular-expression-matching/
# Run: python 0010_regular_expression_matching.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        @cache
        def dp(i, j):
            if j == len(p):
                return i == len(s)
            first = i < len(s) and p[j] in (s[i], ".")
            if j + 1 < len(p) and p[j + 1] == "*":
                return dp(i, j + 2) or (first and dp(i + 1, j))
            return first and dp(i + 1, j + 1)

        return dp(0, 0)


if __name__ == "__main__":
    assert Solution().isMatch("aa", "a") is False
    assert Solution().isMatch("aa", "a*") is True
    assert Solution().isMatch("ab", ".*") is True
    assert Solution().isMatch("aab", "c*a*b") is True
    assert Solution().isMatch("mississippi", "mis*is*p*.") is False
    print("OK")
