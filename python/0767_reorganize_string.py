# 767. Reorganize String (Medium)
# https://leetcode.com/problems/reorganize-string/
# Run: python 0767_reorganize_string.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def reorganizeString(self, s: str) -> str:
        h = [(-c, ch) for ch, c in Counter(s).items()]
        heapify(h)
        res, prev = [], None
        while h:
            c, ch = heappop(h)
            res.append(ch)
            if prev:
                heappush(h, prev)
            prev = (c + 1, ch) if c + 1 < 0 else None
        return "".join(res) if not prev else ""


if __name__ == "__main__":
    r = Solution().reorganizeString("aab")
    assert r == "aba"
    assert Solution().reorganizeString("aaab") == ""
    print("OK")
