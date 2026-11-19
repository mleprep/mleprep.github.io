# 2390. Removing Stars From a String (Medium)
# https://leetcode.com/problems/removing-stars-from-a-string/
# Run: python 2390_removing_stars_from_a_string.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def removeStars(self, s: str) -> str:
        st = []
        for ch in s:
            if ch == "*":
                st.pop()
            else:
                st.append(ch)
        return "".join(st)


if __name__ == "__main__":
    assert Solution().removeStars("leet**cod*e") == "lecoe"
    assert Solution().removeStars("erase*****") == ""
    assert Solution().removeStars("abc") == "abc"
    print("OK")
