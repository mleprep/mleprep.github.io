# 344. Reverse String (Easy)
# https://leetcode.com/problems/reverse-string/
# Run: python 0344_reverse_string.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def reverseString(self, s: List[str]) -> None:
        l, r = 0, len(s) - 1
        while l < r:
            s[l], s[r] = s[r], s[l]
            l, r = l + 1, r - 1


if __name__ == "__main__":
    a = list("hello"); Solution().reverseString(a)
    assert a == list("olleh")
    a = list("Hannah"); Solution().reverseString(a)
    assert a == list("hannaH")
    a = ["x"]; Solution().reverseString(a)
    assert a == ["x"]
    print("OK")
