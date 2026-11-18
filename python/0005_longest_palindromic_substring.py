# 5. Longest Palindromic Substring (Medium)
# https://leetcode.com/problems/longest-palindromic-substring/
# Run: python 0005_longest_palindromic_substring.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def longestPalindrome(self, s: str) -> str:
        best = (0, 0)
        for c in range(2 * len(s) - 1):
            l, r = c // 2, (c + 1) // 2
            while l >= 0 and r < len(s) and s[l] == s[r]:
                l, r = l - 1, r + 1
            if r - l - 1 > best[1] - best[0]:
                best = (l + 1, r)
        return s[best[0]:best[1]]


if __name__ == "__main__":
    assert Solution().longestPalindrome("babad") in ("bab", "aba")
    assert Solution().longestPalindrome("cbbd") == "bb"
    assert Solution().longestPalindrome("a") == "a"
    print("OK")
