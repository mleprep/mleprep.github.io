# 13. Roman to Integer (Easy)
# https://leetcode.com/problems/roman-to-integer/
# Run: python 0013_roman_to_integer.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def romanToInt(self, s: str) -> int:
        val = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
        total = 0
        for i, ch in enumerate(s):
            v = val[ch]
            total += -v if i + 1 < len(s) and val[s[i + 1]] > v else v
        return total


if __name__ == "__main__":
    assert Solution().romanToInt("III") == 3
    assert Solution().romanToInt("LVIII") == 58
    assert Solution().romanToInt("MCMXCIV") == 1994
    print("OK")
