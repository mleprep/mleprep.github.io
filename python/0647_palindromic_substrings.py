# 647. Palindromic Substrings (Medium)
# https://leetcode.com/problems/palindromic-substrings/
# Run: python 0647_palindromic_substrings.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def countSubstrings(self, s: str) -> int:
        count = 0
        for c in range(2 * len(s) - 1):
            l, r = c // 2, (c + 1) // 2
            while l >= 0 and r < len(s) and s[l] == s[r]:
                count += 1
                l, r = l - 1, r + 1
        return count


if __name__ == "__main__":
    assert Solution().countSubstrings("abc") == 3
    assert Solution().countSubstrings("aaa") == 6
    print("OK")
