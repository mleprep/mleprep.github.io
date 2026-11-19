# 290. Word Pattern (Easy)
# https://leetcode.com/problems/word-pattern/
# Run: python 0290_word_pattern.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def wordPattern(self, pattern: str, s: str) -> bool:
        words = s.split()
        if len(words) != len(pattern):
            return False
        p2w, w2p = {}, {}
        for p, w in zip(pattern, words):
            if p2w.setdefault(p, w) != w or w2p.setdefault(w, p) != p:
                return False
        return True


if __name__ == "__main__":
    assert Solution().wordPattern("abba", "dog cat cat dog") is True
    assert Solution().wordPattern("abba", "dog cat cat fish") is False
    assert Solution().wordPattern("aaaa", "dog cat cat dog") is False
    assert Solution().wordPattern("abba", "dog dog dog dog") is False
    assert Solution().wordPattern("aaa", "aa aa aa aa") is False
    print("OK")
