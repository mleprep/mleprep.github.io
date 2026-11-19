# 402. Remove K Digits (Medium)
# https://leetcode.com/problems/remove-k-digits/
# Run: python 0402_remove_k_digits.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def removeKdigits(self, num: str, k: int) -> str:
        st = []
        for d in num:
            while k and st and st[-1] > d:
                st.pop()
                k -= 1
            st.append(d)
        if k:
            st = st[:-k]
        return "".join(st).lstrip("0") or "0"


if __name__ == "__main__":
    assert Solution().removeKdigits("1432219", 3) == "1219"
    assert Solution().removeKdigits("10200", 1) == "200"
    assert Solution().removeKdigits("10", 2) == "0"
    print("OK")
