# 678. Valid Parenthesis String (Medium)
# https://leetcode.com/problems/valid-parenthesis-string/
# Run: python 0678_valid_parenthesis_string.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def checkValidString(self, s: str) -> bool:
        lo = hi = 0
        for ch in s:
            if ch == "(":
                lo, hi = lo + 1, hi + 1
            elif ch == ")":
                lo, hi = lo - 1, hi - 1
            else:
                lo, hi = lo - 1, hi + 1
            if hi < 0:
                return False
            lo = max(lo, 0)
        return lo == 0


if __name__ == "__main__":
    assert Solution().checkValidString("()") is True
    assert Solution().checkValidString("(*)") is True
    assert Solution().checkValidString("(*))") is True
    assert Solution().checkValidString("(((*)") is False
    print("OK")
