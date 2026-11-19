# 745. Prefix and Suffix Search (Hard)
# https://leetcode.com/problems/prefix-and-suffix-search/
# Run: python 0745_prefix_and_suffix_search.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class WordFilter:
    def __init__(self, words: List[str]):
        self.idx = {}
        for i, w in enumerate(words):
            for p in range(len(w) + 1):
                for s in range(len(w) + 1):
                    self.idx[w[:p] + "#" + w[s:]] = i

    def f(self, pref: str, suff: str) -> int:
        return self.idx.get(pref + "#" + suff, -1)


if __name__ == "__main__":
    wf = WordFilter(["apple", "ape", "maple"])
    assert wf.f("a", "e") == 1 and wf.f("ap", "le") == 0 and wf.f("", "le") == 2 and wf.f("b", "") == -1
    assert WordFilter(["apple"]).f("a", "e") == 0
    assert WordFilter(["abc", "abc"]).f("a", "c") == 1
    print("OK")
