# 844. Backspace String Compare (Easy)
# https://leetcode.com/problems/backspace-string-compare/
# Run: python 0844_backspace_string_compare.py
from typing import *
import bisect, collections, functools, heapq, itertools, math, random, string
from collections import Counter, OrderedDict, defaultdict, deque
from functools import cache, lru_cache, reduce
from heapq import heapify, heappop, heappush
from lc import *  # ListNode, TreeNode, make_list, list_vals, make_tree, tree_vals, find_node

class Solution:
    def backspaceCompare(self, s: str, t: str) -> bool:
        def survivors(x):                  # surviving characters, right to left
            skip = 0
            for ch in reversed(x):
                if ch == "#":
                    skip += 1
                elif skip:
                    skip -= 1
                else:
                    yield ch

        return all(a == b for a, b in itertools.zip_longest(survivors(s), survivors(t)))


if __name__ == "__main__":
    assert Solution().backspaceCompare("ab#c", "ad#c") is True
    assert Solution().backspaceCompare("ab##", "c#d#") is True
    assert Solution().backspaceCompare("a#c", "b") is False
    assert Solution().backspaceCompare("a##c", "#a#c") is True
    assert Solution().backspaceCompare("bxj##tw", "bxo#j##tw") is True
    assert Solution().backspaceCompare("ab", "b") is False
    print("OK")
