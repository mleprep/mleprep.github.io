# 392. Is Subsequence (Easy)
# https://leetcode.com/problems/is-subsequence/
# Run: python 0392_is_subsequence.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        i = 0
        for ch in t:
            if i < len(s) and s[i] == ch:
                i += 1
        return i == len(s)


if __name__ == "__main__":
    assert Solution().isSubsequence("abc", "ahbgdc") is True
    assert Solution().isSubsequence("axc", "ahbgdc") is False
    print("OK")
