# 205. Isomorphic Strings (Easy)
# https://leetcode.com/problems/isomorphic-strings/
# Run: python 0205_isomorphic_strings.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        st, ts = {}, {}
        for a, b in zip(s, t):
            if st.setdefault(a, b) != b or ts.setdefault(b, a) != a:
                return False
        return True


if __name__ == "__main__":
    assert Solution().isIsomorphic("egg", "add") is True
    assert Solution().isIsomorphic("foo", "bar") is False
    assert Solution().isIsomorphic("badc", "baba") is False
    print("OK")
