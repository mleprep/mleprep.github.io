# 1657. Determine if Two Strings Are Close (Medium)
# https://leetcode.com/problems/determine-if-two-strings-are-close/
# Run: python 1657_determine_if_two_strings_are_close.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def closeStrings(self, word1: str, word2: str) -> bool:
        c1, c2 = Counter(word1), Counter(word2)
        return c1.keys() == c2.keys() and sorted(c1.values()) == sorted(c2.values())


if __name__ == "__main__":
    assert Solution().closeStrings("abc", "bca") is True
    assert Solution().closeStrings("a", "aa") is False
    assert Solution().closeStrings("cabbba", "abbccc") is True
    assert Solution().closeStrings("uau", "ssx") is False
    print("OK")
